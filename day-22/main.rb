require 'benchmark'

class Game
  attr_reader :boss, :player, :spell_log, :timers, :verbose, :hard

  SPELLBOOK = { 'Magic Missile' => { mana: 53,   damage: 4, duration: 0, hp: 0, recharge: 0,    armour: 0 },
                'Drain' =>         { mana: 73,   damage: 2, duration: 0, hp: 2, recharge: 0,    armour: 0 },
                'Shield' =>        { mana: 113,  damage: 0, duration: 6, hp: 0, recharge: 0,    armour: 7 },
                'Poison' =>        { mana: 173,  damage: 3, duration: 6, hp: 0, recharge: 0,    armour: 0 },
                'Recharge' =>      { mana: 229,  damage: 0, duration: 5, hp: 0, recharge: 101,  armour: 0 } }

  SPELLS = SPELLBOOK.keys

  def initialize(player:, boss:, verbose: false, hard: false)
    @verbose = verbose
    @hard = hard

    @boss = boss.dup
    @player = player.dup
    @timers = SPELLS.map { |spell| [spell, 0] }.to_h

    @spell_log = []
  end

  def mana_spent
    spell_log.inject(0) { |sum, spell| sum += SPELLBOOK[spell][:mana] }
  end

  def cast_spell(spell)
    raise "Cannot cast spell #{spell}!" unless available_spells.include?(spell)

    player[:mana] -= SPELLBOOK[spell][:mana]
    damage = SPELLBOOK[spell][:duration].zero? ? SPELLBOOK[spell][:damage] : 0
    boss[:hp] -= damage
    player[:hp] += SPELLBOOK[spell][:hp]
    timers[spell] = SPELLBOOK[spell][:duration]
    spell_log << spell
    log "Player casts #{spell}, dealing #{damage} damage."
  end

  def available_spells
    inactive_spells = timers.select { |spell, duration| duration < 2 }.keys
    inactive_spells.select { |spell| SPELLBOOK[spell][:mana] <= player[:mana] }
  end

  def player_armour
    timers.select { |spell, duration| duration > 0 }.keys.inject(0) { |sum, spell| sum += SPELLBOOK[spell][:armour] }
  end

  def print_status
    log "- Player has #{player[:hp]} hit points, #{player_armour} armor, #{player[:mana]} mana"
    log "- Boss has #{boss[:hp]} hit points"
  end

  def receive_damage
    damage = [boss[:damage] - player_armour, 1].max
    log "Boss attacks for #{damage} damage."

    player[:hp] -= damage
  end

  def inactive_spell?(spell)
    timers[spell].to_i.zero?
  end

  def activate_spells
    SPELLS.each do |spell|
      next if inactive_spell?(spell)

      timers[spell] = [timers[spell] -1, 0].max
      player[:hp] += SPELLBOOK[spell][:hp]
      player[:mana] += SPELLBOOK[spell][:recharge]
      boss[:hp] -= SPELLBOOK[spell][:damage]

      log "#{spell} deals #{SPELLBOOK[spell][:damage]} damage, it's timer is now #{timers[spell]}."
    end
  end

  def tick(spell)
    player[:hp] -= 1 if hard

    if game_over?
      log "Player dead: #{player_dead?}."
      log "Boss dead: #{boss_dead?}."
      return []
    end

    log '-- Player turn --'
    print_status
    activate_spells
    if boss_dead?
      log 'Boss is dead!'
      return []
    end

    cast_spell spell
    log

    log '-- Boss turn --'
    print_status
    activate_spells

    if boss_dead?
      log 'Boss is dead!'
      return []
    end
    receive_damage

    log
    available_spells
  end

  def status_after(spells)
    spells.each { |spell| tick spell }

    available = (player_dead? || boss_dead?) ? [] : available_spells
    { available_spells: available, mana_spent: mana_spent, player_dead: player_dead?, boss_dead: boss_dead?, spell_log: spell_log }
  end

  def player_dead?
    player[:hp] <= 0 || available_spells.count.zero?
  end

  def boss_dead?
    boss[:hp] <= 0
  end

  def game_over?
    player_dead? || boss_dead?
  end

  def log(message = nil)
    return unless verbose
    puts message
  end
end

###

def chain(spells, hsh, hard: false)
  chains = []

  boss = { hp: 55, damage: 8 }
  player = { hp: 50, mana: 500 }

  g = Game.new player: player, boss: boss, hard: hard

  status = g.status_after spells
  new_spells = status[:available_spells]

  if !status[:player_dead] && status[:boss_dead] && (status[:mana_spent] <= hsh[:best])
    hsh[:best] = status[:mana_spent]
    puts hsh[:best]

    hsh[:spells_logs] << status[:spell_log]
  end
  return [spells] if new_spells.empty? || status[:mana_spent] > hsh[:best] || status[:spell_log].count > hsh[:max_chain]

  new_spells.each do |new_spell|
    sp = spells.dup
    sp << new_spell
    chain(sp, hsh, hard: hard).each { |ch| chains << ch }
  end

  chains
end

###

best = { best: 5000, spells_logs: [], max_chain: 20 }
# time_hard = Benchmark.realtime do
#   ['Poison', 'Shield', 'Recharge', 'Magic Missile', 'Drain'].flat_map { |sp| chain [sp], best, hard: true }
# end
# puts best[:best]
# puts best[:spells_logs].inspect
# puts time_hard

# best[:max_chain] = best[:spells_logs].map(&:length).max
# best[:spells_logs] = []

time_easy = Benchmark.realtime do
  ['Poison', 'Shield', 'Recharge', 'Magic Missile', 'Drain'].flat_map { |sp| chain [sp], best, hard: false }
end
puts best[:best]
puts best[:spells_logs].inspect
puts time_easy

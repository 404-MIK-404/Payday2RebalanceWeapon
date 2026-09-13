Hooks:PostHook( WeaponTweakData, "init", "Rebalance_template", function(self)  

--displayed stats

self.${name}.AMMO_MAX = ${AMMO_MAX}
self.${name}.stats.damage = ${damage} 

self.${name}.fire_mode_data = {
    fire_rate = ${fire_mode_data_fire_rate}
}

self.${name}.auto = {
    fire_rate = ${auto_fire_rate}
}

self.${name}.stats.spread = ${spread}
self.${name}.CLIP_AMMO_MAX = ${CLIP_AMMO_MAX}

end )

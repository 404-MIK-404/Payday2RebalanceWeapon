# PAYDAY 2 Rebalance Weapon

CLI tool for customizing and balancing weapon stats in **PAYDAY 2**.

This tool is based on the [original Steam guide](https://steamcommunity.com/sharedfiles/filedetails/?id=1380067875).

The guide explains how weapon stats work and how to edit them manually. It can also be used as a reference for finding weapon IDs and understanding the available weapon properties.

For a more detailed explanation of the editing process, please refer to the original guide:

[PAYDAY 2 Weapon Stats Guide](https://steamcommunity.com/sharedfiles/filedetails/?id=1380067875)

### Help

Run the following command to display the available options:

```bash
./Payday2RebalanceWeapon.exe -h
```

#### Output

```text
Usage of Payday2RebalanceWeapon.exe:
  -a string
        Maximum ammo reserve (default "200")
  -auto_fr string
        Automatic fire rate value (60 / fire_rate) (default "0.071")
  -clip_ammo string
        Maximum magazine capacity (default "30")
  -d string
        Weapon damage (default "85")
  -f string
        Folder name (required) (default "amcar_weapon_template")
  -fire_mode_fr string
        Fire rate value. RPM is calculated as 60 / fire_rate.
        For example, 600 RPM = 0.1, 850 RPM = 0.07058, and so on.
        (default "0.071")
  -n string
        Weapon ID (required; check weapon IDs at https://steamcommunity.com/sharedfiles/filedetails/?id=1380067875)
        (default "amcar")
  -s string
        Weapon spread (default "13")
```

## Required Parameters

### `-f` — Folder Name

Specifies the folder containing the weapon template.

```bash
-f string
```

**Default:** `amcar_weapon_template`

### `-n` — Weapon ID

Specifies the weapon ID to modify.

```bash
-n string
```

**Default:** `amcar`

You can find the available weapon IDs in the [Steam guide](https://steamcommunity.com/sharedfiles/filedetails/?id=1380067875).

## Options

### `-a` — Maximum Ammo Reserve

Sets the maximum ammo reserve.

```bash
-a string
```

**Default:** `200`

The reserve ammo value is applied to the `AMMO_MAX` section, following the same process as the damage value.

### `-auto_fr` — Automatic Fire Rate

Sets the automatic fire rate value.

```bash
-auto_fr string
```

**Default:** `0.071`

This option works with the `-fire_mode_fr` parameter.

### `-clip_ammo` — Magazine Capacity

Sets the maximum magazine capacity.

```bash
-clip_ammo string
```

**Default:** `30`

The magazine size is applied to the `CLIP_AMMO` section. This value is written as two numbers in decimal format:

* The first number represents the base ammo pickup.
* The second number represents the lower pickup chance, unless **Fully Loaded** is used.

### `-d` — Weapon Damage

Sets the weapon damage.

```bash
-d string
```

**Default:** `85`

### `-fire_mode_fr` — Fire Rate

Sets the weapon's fire rate value.

```bash
-fire_mode_fr string
```

**Default:** `0.071`

The fire rate is calculated using the following formula:

```text
RPM = 60 / fire_rate
```

For example:

| Fire rate value | Approximate RPM |
| --------------- | --------------- |
| `0.1`           | 600 RPM         |
| `0.07058`       | 850 RPM         |

**Note:** Lower values result in a higher fire rate, while higher values result in a slower fire rate. You may need to adjust this value several times to achieve the desired result.

### `-s` — Weapon Spread

Sets the weapon's spread value.

```bash
-s string
```

**Default:** `13`

The spread value is applied to the `.spread` section.

**Note:** The maximum spread value you can set is `30`.


## Demo

### Example Command

Example of running the tool:

```bash
Payday2RebalanceWeapon.exe -n ak47 -f ak47_great_again -a 400
```

<img width="835" height="103" alt="изображение" src="https://github.com/user-attachments/assets/abc24cbb-56db-4036-b4d9-1fb18fa4bd72" />

### Generated Folder Structure

After running the tool, the following folder and files are generated:

- desktop.ini
- mod.txt
- Weapon_Template_Rebalance.lua

<img width="628" height="151" alt="изображение" src="https://github.com/user-attachments/assets/b4c1811c-9737-4cf1-8dc1-20b1aac81245" />

### Generated Lua File

The generated Weapon_Template_Rebalance.lua contains the modified weapon stats:

<img width="916" height="475" alt="изображение" src="https://github.com/user-attachments/assets/d8e95188-8a2c-47b3-9bcb-6a77261dbe09" />


## Credits

Weapon stat explanations and the original editing guide are based on the work of the original author.

[Original Steam Guide](https://steamcommunity.com/sharedfiles/filedetails/?id=1380067875)

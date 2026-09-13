package main

import (
	"embed"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

var globalFiles = []string{
	"desktop.ini",
	"mod.txt",
	"Weapon_Template_Rebalance.lua",
}

//go:embed template/*
var templateFS embed.FS

func main() {

	folder := flag.String(
		"f",
		"amcar_weapon_template",
		"Folder name (required)",
	)

	name := flag.String(
		"n",
		"amcar",
		"Weapon ID (required; check weapon ID's at https://steamcommunity.com/sharedfiles/filedetails/?id=1380067875)",
	)

	ammoMax := flag.String(
		"a",
		"200",
		"Maximum ammo reserve",
	)

	damage := flag.String(
		"d",
		"85",
		"Weapon damage",
	)

	fireModDataFireRate := flag.String(
		"fire_mode_fr",
		"0.071",
		"Fire rate value. RPM is calculated as 60 / fire_rate. For example, 600 RPM = 0.1, 850 RPM = 0.07058, and so on.",
	)

	autoFireRate := flag.String(
		"auto_fr",
		"0.071",
		"Automatic fire rate value (60 / fire_rate)",
	)

	spread := flag.String(
		"s",
		"13",
		"Weapon spread",
	)

	clipAmmoMax := flag.String(
		"clip_ammo",
		"30",
		"Maximum magazine capacity",
	)

	flag.Parse()

	validateRequiredFields(*name, *folder)

	values := map[string]string{
		"name":                     *name,
		"AMMO_MAX":                 *ammoMax,
		"damage":                   *damage,
		"fire_mode_data_fire_rate": *fireModDataFireRate,
		"auto_fire_rate":           *autoFireRate,
		"spread":                   *spread,
		"CLIP_AMMO_MAX":            *clipAmmoMax,
	}

	createFolder(*folder)

	for _, f := range globalFiles {
		dst := filepath.Join(*folder, f)

		err := extractFile(f, dst)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Printf("File extracted: %s\n", dst)

		err = renderTemplate(dst, values)
		if err != nil {
			fmt.Println("Template rendering error:", err)
			continue
		}
	}
	fmt.Printf("Template %s created successfully!\n", *folder)
}

func validateRequiredFields(name string, folder string) {
	if name == "" || folder == "" {
		fmt.Println("Error: required flags -n (weapon ID) and -f (folder name) must be provided.")
		flag.Usage()
		os.Exit(1)
	}
}

func extractFile(name string, dst string) error {
	data, err := templateFS.ReadFile("template/" + name)
	if err != nil {
		return fmt.Errorf("failed to read embedded file %q: %w", name, err)
	}

	if err := os.WriteFile(dst, data, 0644); err != nil {
		return fmt.Errorf("failed to write file %q: %w", dst, err)
	}
	return nil
}

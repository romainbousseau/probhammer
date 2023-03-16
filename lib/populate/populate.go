package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/romainbousseau/probhammer/internal/models"
	"github.com/romainbousseau/probhammer/internal/utils"
	"gorm.io/gorm"
)

// type populatable interface {
// 	models.Datasheet | models.Faction | models.DatasheetUnit
// }

func main() {
	// Load Config
	config, err := utils.LoadConfig(".", ".env")
	if err != nil {
		log.Fatal("unable to load config:", err)
	}
	fmt.Println("- config loaded")

	// Open DB connection
	db, err := utils.OpenDBConnection(config)
	if err != nil {
		log.Fatal("unable to open db connection:", err)
	}
	fmt.Println("- connected to DB")

	// Drop tables
	err = utils.DropAllTables(db)
	if err != nil {
		log.Fatal("unable to drop tables", err)
	}
	fmt.Println("- tables dropped")

	// Migrate
	err = utils.Migrate(db)
	if err != nil {
		log.Fatal("unable to migrates:", err)
	}
	fmt.Println("- migration completed")

	// Set reader
	// Allows use pipe as delimiter and escape double quotes
	gocsv.SetCSVReader(func(in io.Reader) gocsv.CSVReader {
		r := csv.NewReader(in)
		r.Comma = '|'
		r.LazyQuotes = true
		return r
	})

	// Factions
	err = createFactions(db)
	if err != nil {
		log.Fatal("unable create factions:", err)
	}
	fmt.Println("- factions created")

	// Datasheets
	err = createDatasheets(db)
	if err != nil {
		log.Fatal("unable create datasheets:", err)
	}
	fmt.Println("- datasheets created")

	// Units
	err = createUnits(db)
	if err != nil {
		log.Fatal("unable create units:", err)
	}
	fmt.Println("- units created")
}

func createFactions(db *gorm.DB) error {
	file, err := os.Open("Factions.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	factions := []*models.Faction{}

	if err := gocsv.UnmarshalFile(file, &factions); err != nil {
		return err
	}

	for _, faction := range factions {
		err = db.Create(faction).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func createDatasheets(db *gorm.DB) error {
	file, err := os.Open("Datasheets.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '|'
	reader.LazyQuotes = true

	datasheets := []*models.Datasheet{}

	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records[1:] {
		iD, err := strconv.ParseUint(strings.TrimLeft(record[0], "0"), 10, 32)
		if err != nil {
			return err
		}

		datasheets = append(datasheets, &models.Datasheet{
			ID: uint(iD),
			Name: record[1],
			FactionId: record[3],
		})
	}

	for _, datasheet := range datasheets {
		err = db.Create(datasheet).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func createUnits(db *gorm.DB) error {
	file, err := os.Open("Units.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.Comma = '|'
	reader.LazyQuotes = true

	units := []*models.Unit{}

	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range records[1:] {
		datasheetID, err := strconv.ParseUint(strings.TrimLeft(record[0], "0"), 10, 32)
		if err != nil {
			return err
		}

		position, err := strconv.Atoi(record[1])
		if err != nil {
			return err
		}

		wSkill, err := strconv.Atoi(trimRightPlus(removeDash(record[4])))
		if err != nil {
			return err
		}

		bSkill, err := strconv.Atoi(trimRightPlus(removeDash(record[5])))
		if err != nil {
			return err
		}

		strength, err := strconv.Atoi(removeDash(record[6]))
		if err != nil {
			return err
		}

		toughness, err := strconv.Atoi(removeDash(record[7]))
		if err != nil {
			return err
		}

		wounds, err := strconv.Atoi(removeDash(record[8]))
		if err != nil {
			return err
		}

		attacks := removeDash(record[9])

		var save int
		var iSave int
		if strings.Contains(record[11], "/") {
			split := strings.Split(record[11], "/")
			save, err = strconv.Atoi(trimRightPlus(removeDash(split[0])))
			if err != nil {
				return err
			}
			iSave, err = strconv.Atoi(trimRightPlus(removeDash(split[1])))
			if err != nil {
				return err
			}
		} else {
			save, err = strconv.Atoi(trimRightPlus(removeDash(record[11])))
			if err != nil {
				return err
			}
			iSave = 0
		}

		units = append(units, &models.Unit{
			DatasheetID:      uint(datasheetID),
			Position:         int32(position),
			Name:             record[2],
			WeaponSkill:      int32(wSkill),
			BallisticSkill:   int32(bSkill),
			Strength:         int32(strength),
			Toughness:        int32(toughness),
			Wounds:           int32(wounds),
			Attacks:          attacks,
			Save:             int32(save),
			InvulnerableSave: int32(iSave),
		})
	}

	for _, unit := range units {
		err = db.Create(unit).Error
		if err != nil {
			return err
		}
	}

	return nil
}

func removeDash(str string) string {
	return strings.Replace(str, "-", "0", 1)
}

func trimRightPlus(str string) string {
	return strings.TrimRight(str, "+")
}

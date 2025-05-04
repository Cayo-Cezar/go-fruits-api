package schedulers

import (
	"log"
	"time"

	"github.com/Cayo-Cezar/go-fruits-api/services"
	"github.com/robfig/cron/v3"
)

func StartCron(fruitService *services.FruitService) {
	c := cron.New(cron.WithLocation(time.FixedZone("America/Sao_Paulo", -3*60*60)))

	_, err := c.AddFunc("0 0 * * *", func() {
		log.Println("🕛 Executando crawler (agendado)...")
		if err := fruitService.FetchAndSaveFruits(); err != nil {
			log.Println("Erro ao importar frutas:", err)
		}
	})

	if err != nil {
		log.Println("Erro ao agendar cron:", err)
	}

	c.Start()
	log.Println("Cron iniciado — agendamento configurado")
}

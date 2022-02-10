package logic

import "github.com/rs/zerolog/log"

var pastScores = []float64{100.0, 1.0}

func GetSize() float64 {
	latestScore := pastScores[len(pastScores)-1]
	var totalScore float64
	for _, score := range pastScores {
		totalScore += float64(score)
	}
	avgScore := totalScore / float64(len(pastScores))
	performance := avgScore - latestScore

	var size float64
	if performance > 0.0 {
		if s := 100.0 + performance*50.0; s < 4000.0 {
			size = s
		} else {
			size = 4000.0
		}
	} else if performance <= 0.0 {
		size = 100.0 + performance*20.0
		if size < 0.0 {
			size = 1.0
		}
	}
	log.Info().Msgf("performance: [%v]; size: [%v]\n", performance, size)
	return size
}

func SetScore(newScore float64) bool {
	pastScores = append(pastScores, newScore)
	return true
}

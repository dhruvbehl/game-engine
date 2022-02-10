package logic


var pastScores = []float64{5.0, 4.0, 3.0, 2.0, 1.0}

func GetSize() float64 {
	latestScore := pastScores[len(pastScores) - 1]
	var totalScore float64
	for score, _ := range pastScores {
		totalScore += float64(score)
	}
	avgScore := totalScore / float64(len(pastScores))
	performance := latestScore - avgScore
	if performance > 0.0 {
		if size := 500.0 + performance * 50.0; size < 2000.0 {
			return size
		} else {
			return 2000.0
		}
	}
	if performance > -0.5 && performance <= 0.0 {
		return 100.0 + performance * 20.0
	}
	return 10.0
}

func SetScore(newScore float64) bool {
	pastScores = append(pastScores, newScore)
	return true
}
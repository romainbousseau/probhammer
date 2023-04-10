package calculator

import (
	"testing"

	"gopkg.in/go-playground/assert.v1"
)

func Test_Roll(t *testing.T) {
	t.Run("Returns the correct results", func(t *testing.T) {
		assert.Equal(t, Roll(4, 2), Result{Success: 3, Failure: 1, Prob: 0.8333333333333334})
		assert.Equal(t, Roll(4, 3), Result{Success: 3, Failure: 1, Prob: 0.6666666666666666})
		assert.Equal(t, Roll(4, 4), Result{Success: 2, Failure: 2, Prob: 0.5})
		assert.Equal(t, Roll(4, 5), Result{Success: 1, Failure: 3, Prob: 0.3333333333333333})
		assert.Equal(t, Roll(4, 6), Result{Success: 1, Failure: 3, Prob: 0.16666666666666666})

		assert.Equal(t, Roll(10, 2), Result{Success: 8, Failure: 2, Prob: 0.8333333333333334})
		assert.Equal(t, Roll(10, 3), Result{Success: 7, Failure: 3, Prob: 0.6666666666666666})
		assert.Equal(t, Roll(10, 4), Result{Success: 5, Failure: 5, Prob: 0.5})
		assert.Equal(t, Roll(10, 5), Result{Success: 3, Failure: 7, Prob: 0.3333333333333333})
		assert.Equal(t, Roll(10, 6), Result{Success: 2, Failure: 8, Prob: 0.16666666666666666})
	})
}

func Test_Calculate(t *testing.T) {
	t.Run("Returns the correct results with superior toughness", func(t *testing.T) {
		c := Calculator{
			Attacks:   10,
			Damage: 	 2,
			HitSkill:  3,
			Strength:  4,
			Toughness: 3,
			Save:      5,
		}

		c.Calculate()

		assert.Equal(t, c.Results.Hit, Result{Success: 7, Failure: 3, Prob: 0.6666666666666666})
		assert.Equal(t, c.Results.Wound, Result{Success: 5, Failure: 2, Prob: 0.6666666666666666})
		assert.Equal(t, c.Results.Save, Result{Success: 2, Failure: 3, Prob: 0.3333333333333333})
		assert.Equal(t, c.Results.ExpectedDamage, int32(6))
	})

	t.Run("Returns the correct results with twice superior toughness", func(t *testing.T) {
		c := Calculator{
			Attacks:   4,
			HitSkill:  2,
			Strength:  8,
			Toughness: 4,
			Save:      3,
		}

		c.Calculate()

		assert.Equal(t, c.Results.Hit, Result{Success: 3, Failure: 1, Prob: 0.8333333333333334})
		assert.Equal(t, c.Results.Wound, Result{Success: 2, Failure: 1, Prob: 0.8333333333333334})
		assert.Equal(t, c.Results.Save, Result{Success: 1, Failure: 1, Prob: 0.6666666666666666})
	})

	t.Run("Returns the correct results with equal toughness", func(t *testing.T) {
		c := Calculator{
			Attacks:   8,
			HitSkill:  4,
			Strength:  4,
			Toughness: 4,
			Save:      4,
		}

		c.Calculate()

		assert.Equal(t, c.Results.Hit, Result{Success: 4, Failure: 4, Prob: 0.5})
		assert.Equal(t, c.Results.Wound, Result{Success: 2, Failure: 2, Prob: 0.5})
		assert.Equal(t, c.Results.Save, Result{Success: 1, Failure: 1, Prob: 0.5})
	})

	t.Run("Returns the correct results with inferior toughness", func(t *testing.T) {
		c := Calculator{
			Attacks:   20,
			HitSkill:  5,
			Strength:  3,
			Toughness: 5,
			Save:      2,
		}

		c.Calculate()

		assert.Equal(t, c.Results.Hit, Result{Success: 7, Failure: 13, Prob: 0.3333333333333333})
		assert.Equal(t, c.Results.Wound, Result{Success: 2, Failure: 5, Prob: 0.3333333333333333})
		assert.Equal(t, c.Results.Save, Result{Success: 2, Failure: 0, Prob: 0.8333333333333334})
	})

	t.Run("Returns the correct results with twice inferior toughness", func(t *testing.T) {
		c := Calculator{
			Attacks:   2,
			HitSkill:  3,
			Strength:  5,
			Toughness: 10,
			Save:      5,
		}

		c.Calculate()

		assert.Equal(t, c.Results.Hit, Result{Success: 1, Failure: 1, Prob: 0.6666666666666666})
		assert.Equal(t, c.Results.Wound, Result{Success: 0, Failure: 1, Prob: 0.3333333333333333})
		assert.Equal(t, c.Results.Save, Result{Success: 0, Failure: 0, Prob: 0.3333333333333333})
	})

	t.Run("Returns the correct results with no hits", func(t *testing.T) {
		c := Calculator{
			Attacks:   1,
			Damage: 	 2,
			HitSkill:  5,
			Strength:  4,
			Toughness: 4,
			Save:      4,
		}

		c.Calculate()

		assert.Equal(t, c.Results.Hit, Result{Success: 0, Failure: 1, Prob: 0.3333333333333333})
		assert.Equal(t, c.Results.Wound, Result{Success: 0, Failure: 0, Prob: 0.5})
		assert.Equal(t, c.Results.Save, Result{Success: 0, Failure: 0, Prob: 0.5})
		assert.Equal(t, c.Results.ExpectedDamage, int32(0))
	})
}

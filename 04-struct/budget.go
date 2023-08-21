package main

import (
	"fmt"
	"time"
)

type Budget struct {
	CampaignID string
	Balance    float64
	Expires    time.Time
}

func NewBudget(campaignID string, balance float64, expires time.Time) (*Budget, error) {
	if campaignID == "" {
		return nil, fmt.Errorf("empty campaignID")
	}

	if balance <= 0 {
		return nil, fmt.Errorf("balance must be bigger than 0")
	}

	b := Budget{
		CampaignID: campaignID,
		Balance:    balance,
		Expires:    expires,
	}

	return &b, nil
}

func (b Budget) TimeLeft() time.Duration {
	return b.Expires.Sub(time.Now().UTC())
}

// if we want to change property of a struct we need to use pointer receiver
func (b *Budget) Update(sum float64) {
	b.Balance += sum
}

func main() {
	b1 := Budget{"Kitten", 22.3, time.Now().Add(7 * 24 * time.Hour)}
	fmt.Println(b1)
	fmt.Println(b1.TimeLeft())

	b1.Update(10.5)
	fmt.Println(b1)

	fmt.Printf("%#v\n", b1)

	fmt.Println(b1.CampaignID)

	b2 := Budget{
		Balance:    19.2,
		CampaignID: "Puppy",
	}

	fmt.Println(b2)

	var b3 Budget

	fmt.Println(b3)

	expire := time.Now().Add(7 * 24 * time.Hour)

	b4, err := NewBudget("Puppy", 32.2, expire)

	if err != nil {
		fmt.Println("ERROR: ", err)
	} else {
		fmt.Printf("%#v\n", b4)
	}

	b5, err := NewBudget("Kitten", -22.2, expire)

	if err != nil {
		fmt.Println("ERROR: ", err)
	} else {
		fmt.Printf("%#v\n", b5)
	}
}

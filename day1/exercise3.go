package main

import "fmt"

const (
	TotalWorkingDays      = 28
	PermanentBasicPay     = 500
	ContractualBasicPay   = 100
	FreelancerRatePerHour = 10
)

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	basicPay int
}

type Contractual struct {
	basicPay int
}

type Freelancer struct {
	ratePerHour int
	totalHours  int
}

func (f *Freelancer) setTotalHours(hours int) {
	f.totalHours = hours
}

func (p Permanent) CalculateSalary() int {
	return p.basicPay * TotalWorkingDays
}

func (c Contractual) CalculateSalary() int {
	return c.basicPay * TotalWorkingDays
}

func (f Freelancer) CalculateSalary() int {
	return f.ratePerHour * f.totalHours
}

func main() {
	pemp := Permanent{
		basicPay: PermanentBasicPay,
	}
	cemp := Contractual{
		basicPay: ContractualBasicPay,
	}
	freelancer := Freelancer{
		ratePerHour: FreelancerRatePerHour,
	}

	var totalWorkingHours int
	fmt.Println("Enter total working hours for freelancer : ")

	_, err := fmt.Scanf("%d", &totalWorkingHours)

	if err != nil {
		fmt.Println(err)
	}
	freelancer.setTotalHours(totalWorkingHours)

	fmt.Println(pemp.CalculateSalary())
	fmt.Println(cemp.CalculateSalary())
	fmt.Println(freelancer.CalculateSalary())

}

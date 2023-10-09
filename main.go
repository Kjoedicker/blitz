package main

func main() {
	printLogo()

	testPlanPath := parseTestPlanPath()
	testPlan := loadTestPlan(testPlanPath)

	testPlan.begin()
}

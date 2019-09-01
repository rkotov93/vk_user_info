package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/rkotov93/easyvk-go/easyvk"
)

const authURL = "https://oauth.vk.com/authorize" +
	"?client_id=6826999" +
	"&scope=" +
	"&redirect_uri=https://oauth.vk.com/blank.html" +
	"&display=wap" +
	"&v=5.63" +
	"&response_type=token"

func main() {
	token, low, upp, id := parseFlags()
	vk := easyvk.WithToken(token)
	user := getUser(&vk, id)
	bdate := defineUserBDate(user, &vk, low, upp)

	data := collectData(user, bdate)
	printData(data)
}

func parseFlags() (string, int, int, string) {
	token := flag.String("t", "", "Access token (optional)")
	low := flag.Int("l", 1, "Lower bound (optional)")
	upp := flag.Int("u", 100, "Upper bound (optional)")
	flag.Parse()

	if *token == "" {
		fmt.Println(
			"Access token was not provided. Use -t flag so specify it.",
			"To get access token proceed to", authURL,
			"login and copy 'access_token' param from URL",
		)
		os.Exit(1)
	}

	if len(flag.Args()) < 1 {
		fmt.Println("ID was not specified.")
		os.Exit(1)
	}
	id := flag.Args()[0]

	return *token, *low, *upp, id
}

func defineUserBDate(user *easyvk.User, vk *easyvk.VK, low int, upp int) string {
	fmt.Print("Calculating user's age...")
	defer fmt.Print("\n\n")

	dateArr := strings.Split(user.Bdate, ".")
	if len(dateArr) < 2 {
		return ""
	}

	bDay, _ := strconv.Atoi(dateArr[0])
	bMonth, _ := strconv.Atoi(dateArr[1])
	var age, bYear int

	if len(dateArr) == 3 {
		bYear, _ = strconv.Atoi(dateArr[2])
		age = calculateAgeFrombYear(bYear, bMonth, bDay)
	} else {
		age = getUserAge(vk, user, low, upp)
		bYear = calculatebYearFromAge(age, bMonth, bDay)
	}

	return fmt.Sprintf("%d.%d.%d (%d years)", bDay, bMonth, bYear, age)
}

func calculatebYearFromAge(age int, bMonth, bDay int) int {
	now := time.Now()
	currentYearBDate := time.Date(now.Year(), time.Month(bMonth), bDay, 0, 0, 0, 0, time.UTC)
	if now.Before(currentYearBDate) {
		age++
	}
	now = now.AddDate(-age, 0, 0)

	return now.Year()
}

func calculateAgeFrombYear(bYear int, bMonth int, bDay int) int {
	now := time.Now()
	age := now.Year() - bYear

	currentYearBDate := time.Date(now.Year(), time.Month(bMonth), bDay, 0, 0, 0, 0, time.UTC)
	if now.Before(currentYearBDate) {
		age--
	}

	return age
}

func collectData(user *easyvk.User, bdate string) map[string]string {
	data := make(map[string]string)

	data["ID"] = strconv.FormatUint(user.ID, 10)
	data["Name"] = user.FirstName + " " + user.LastName
	data["Nickname"] = user.Nickname
	data["Sex"] = defineSex(user.Sex)
	data["BDate"] = bdate
	data["City"] = user.City.Title
	data["Country"] = user.Country.Title
	data["Relation"] = strconv.Itoa(int(user.Relation))

	return data
}

func printData(data map[string]string) {
	fmt.Println("==========")
	fmt.Println("USER DATA:")
	fmt.Println("==========")

	keys := []string{"ID", "Name", "Nickname", "Sex", "BDate", "City", "Country", "Relation"}
	for i := range keys {
		key := keys[i]
		fmt.Println(key + ": " + data[key])
	}

	fmt.Println("==========")
}

func defineSex(sex uint8) string {
	switch sex {
	case 1:
		return "Woman"
	case 2:
		return "Man"
	}

	return "undefined"
}

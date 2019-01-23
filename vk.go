package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/rkotov93/easyvk-go/easyvk"
)

type checkResult struct {
	fits bool
	low  int
	upp  int
}

func getUser(vk *easyvk.VK, id string) *easyvk.User {
	ids := []string{id}
	fields := []string{
		"sex", "nickname", "domain", "city", "country", "home_town",
		"status", "bdate", "interests", "relation",
	}

	fmt.Printf("Getting general available informatin about %#v...\n", id)
	users, err := vk.Users.Get(ids, fields, "nom")
	handleError(err)
	return &users[0]
}

func getUserAge(vk *easyvk.VK, user *easyvk.User, low int, upp int) int {
	dateArr := strings.Split(user.Bdate, ".")

	filters := make(map[string]string)
	filters["city"] = strconv.Itoa(user.City.ID)
	filters["country"] = strconv.Itoa(user.Country.ID)
	filters["hometown"] = user.Hometown
	filters["sex"] = strconv.Itoa(int(user.Sex))
	filters["status"] = strconv.Itoa(int(user.Relation))
	filters["birth_day"] = dateArr[0]
	filters["birth_month"] = dateArr[1]
	filters["interests"] = user.Interests

	return binaryUsersFilter(vk, user, low, upp, filters)
}

func binaryUsersFilter(vk *easyvk.VK, user *easyvk.User, low int, upp int, filters map[string]string) int {
	if low == upp {
		return low
	}

	mid := (upp-low)/2 + low
	fits := make(chan checkResult)

	go checkUserFits(vk, user, low, mid, filters, fits)
	go checkUserFits(vk, user, mid+1, upp, filters, fits)

	var result checkResult
	for i := 0; i < 2; i++ {
		tmp := <-fits
		if tmp.fits {
			result = tmp
		}
	}

	return binaryUsersFilter(vk, user, result.low, result.upp, filters)
}

func checkUserFits(vk *easyvk.VK, user *easyvk.User, low int, upp int, p map[string]string, fits chan checkResult) {
	time.Sleep(time.Second)
	fmt.Print(".")

	q := user.FirstName + " " + user.LastName
	params := copyParams(p)
	params["age_from"] = strconv.Itoa(low)
	params["age_to"] = strconv.Itoa(upp)

	filteredResults := getFilteredResults(vk, q, params)
	isUserFits := isFilteredResultsContainUser(filteredResults, user)

	fits <- checkResult{isUserFits, low, upp}
}

func getFilteredResults(vk *easyvk.VK, q string, params map[string]string) *easyvk.UsersSearchResults {
	filteredResults, err := vk.Users.Search(q, params)
	handleError(err)

	return filteredResults
}

func isFilteredResultsContainUser(filteredResults *easyvk.UsersSearchResults, user *easyvk.User) bool {
	users := filteredResults.Items

	index := -1
	for i := 0; i < filteredResults.Count; i++ {
		if users[i].ID == user.ID {
			index = i
		}
	}

	return index != -1
}

func copyParams(params map[string]string) map[string]string {
	newParams := make(map[string]string)
	for k, v := range params {
		newParams[k] = v
	}

	return newParams
}

func handleError(err error) {
	if err != nil {
		fmt.Println("Something went wrong:", err)
		os.Exit(1)
	}
}

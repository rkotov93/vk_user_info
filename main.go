package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	from := flag.Int("f", 1, "Lower bound (optional)")
	to := flag.Int("t", 100, "Upper bound (optional)")
	flag.Parse()

	if len(flag.Args()) < 1 {
		fmt.Println("ID was not specified.")
		os.Exit(1)
	}
	id := flag.Args()[0]

	fmt.Println(id, *from, *to)
}

// city, country, home_town => hometown, sex, status, bdate => (birth_day, birth_month), interests
// city,country,home_town,sex,status,bdate,interests
// photo_id,verified,sex,bdate,city,country,home_town,has_photo,photo_50,photo_100,photo_200_orig,photo_200,photo_400_orig,photo_max,photo_max_orig,online,lists,domain,has_mobile,contacts,site,education,universities,schools,status,last_seen,followers_count,common_count,occupation,nickname,relatives,relation,personal,connections,exports,wall_comments,activities,interests,music,movies,tv,books,games,about,quotes,can_post,can_see_all_posts,can_see_audio,can_write_private_message,can_send_friend_request,is_favorite,is_hidden_from_feed,timezone,screen_name,maiden_name,crop_photo,is_friend,friend_status,career,military,blacklisted,blacklisted_by_me

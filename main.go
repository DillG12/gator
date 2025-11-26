package main

import (
	"log"
	"os"

	"database/sql"

	"github.com/DillG12/gator/internal/config"
	"github.com/DillG12/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		log.Fatalf("Failed to open database: %v", err)
	}
	defer db.Close()

	dbQueries := database.New(db)

	s := &state{cfg: &cfg, db: dbQueries}

	cmds := &commands{list: make(map[string]func(*state, command) error)}
	cmds.register("login", loginHandler)
	cmds.register("register", registerHandler)
	cmds.register("reset", resetHandler)
	cmds.register("users", usersHandler)
	cmds.register("agg", aggHandler)
	cmds.register("addfeed", middlewareLoggedIn(addFeedHandler))
	cmds.register("feeds", feedsHandler)
	cmds.register("follow", middlewareLoggedIn(followHandler))
	cmds.register("following", middlewareLoggedIn(followingHandler))
	cmds.register("unfollow", middlewareLoggedIn(unfollowHandler))
	cmds.register("browse", middlewareLoggedIn(browseHandler))

	args := os.Args[1:]
	if len(args) == 0 {
		log.Fatalf("No command provided")
	}

	cmd := command{name: args[0], args: args[1:]}

	err = cmds.run(s, cmd)
	if err != nil {
		log.Fatalf("Command failed: %v", err)
	}

	// Read back the config to verify changes

	cfg, err = config.Read()
	if err != nil {
		log.Fatalf("Failed to read config: %v", err)
	}

	//fmt.Printf("%v\n", cfg)
}

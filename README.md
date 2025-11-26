# gator
Gator is a blog aggregator cli tool written in golang and SQL. 
In order to run gator you will need to have golang and postgres installed on your computer. You can install [go here](https://go.dev/doc/install) and [postgres here](https://www.postgresql.org/download/).
Once those are installed you can download the gator cli tool by typing go install github.com/DillG12/gator in your command line.
Once that is done you will need to create a config file in your home directory. Manually create a file in your home directory ~/.gatorconfig.json and paste this code in to it.
```
{
  "db_url": "postgres://example"
}
```
Once that is done you should be able to use the gator cli tool! You can register a user by typing ./gator register <username>. Some other commands you can use are:
- login <username> (logs in the username provided)
- reset (deletes all users)
- users (lists all users)
- agg <time_between_requests> (saves posts from current users feeds with a given time between downloads "1m", "10s", etc)
- addfeed <name> <url> (adds a feed to current users followed feeds)
- feeds (lists all feeds and users who are following them)
- follow <url> (saves an existing feed to current users followed feeds)
- following (lists current users followed feeds)
- unfollow <url> (deletes a feed from current users followed feeds)
- browse <number_of_posts *optional*> (returns most recent posts from feeds and links to them. If no number is provided for amount of posts returned, browse command will return 2 posts)

  That should be all you need to know to get started! I hope you enjoy this project. I had a lot of fun making it!

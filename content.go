package main

// UserName is name of the user in portfolio
const UserName = "Abhilash Gnan"

// Blocks holds actual block data
var Blocks = []Block{
	// Projects page content
	&ProjectsBlock{
		Meta: Meta{
			Name:    "projects",
			Heading: "Projects",
		},
		Projects: []Project{
			{
				Title: "Snap",
				Desc:  "Snapshot file testing tool written in bash.",
				Link:  "https://github.com/jan25/snap",
			},
			{
				Title: "Termracer",
				Desc:  "Terminal based touch typing tutor written in Go. Works completely offline.",
				Link:  "https://github.com/jan25/termracer",
			},
			{
				Title: "Move my files",
				Desc:  "Easy to use CLI tool to organize files on your computer written in Python.",
				Link:  "https://github.com/jan25/move-my-files",
			},
			{
				Title: "Changelog gen",
				Desc:  "CLI tool implemented in Elixir to automate generation of CHANGELOG file for Github repositories.",
				Link:  "https://github.com/jan25/changeloggen",
			},
			{
				Title: "Memory test",
				Desc:  "Test your short term memory skills using this simple game written using ReactJS, inspired by a BBC Earth video.",
				Link:  "https://github.com/jan25/memory-test",
			},
			{
				Title: "Skripts",
				Desc:  "Execute node.js scripts inside a web browser and share using unique links. Implemented using Node and Vuejs.",
				Link:  "https://github.com/jan25/skripts",
			},
			{
				Title: "HotRod",
				Desc:  "Jaeger's HotRod app rewritten in Python. Orignal demo app was written in Go.",
				Link:  "https://github.com/jan25/hotrod-python",
			},
			{
				Title: "Sudoku",
				Desc:  "React app with fairly difficult puzzles that are pregenerated using a Python script.",
				Link:  "https://github.com/jan25/sudoku",
			},
			{
				Title: "F1 Lights",
				Desc:  "Test your reaction times with Formula1 lights implemented using Vuejs.",
				Link:  "https://github.com/jan25/f1-lights",
			},
			{
				Title: "Track ISS",
				Desc:  "Web app for real time position tracking of International space station written in React",
				Link:  "https://github.com/jan25/track-iss",
			},
			{
				Title: "Geo GraphQL API",
				Desc: "Powerful recursive Geo taxonomy API. Imagine querying cities of a country and cities in neighboring countries and so on.",
				Link: "https://github.com/jan25/geo-graphql",
			},
			{
				Title: "Shuffle",
				Desc: "Shuffling algorithm written in Typescript, inspired by a Spotify blog post on shuffling playlist songs without needing for true randomness.",
				Link: "https://github.com/jan25/shuffle",
			},
			{
				Title: "Game of life",
				Desc:  "React app showing Gosper's gilder gun, one of infinitely continuing patterns in Game of Life. R.I.P Conway.",
				Link:  "https://github.com/jan25/gameoflife",
			},
			{
				Title: "Valen10",
				Desc:  "Breakout game with emojis and silly gifs implemented on a Feb the 14th. Written in vanilla JS with offline support.",
				Link:  "https://github.com/jan25/sandbox-2020/tree/master/valen10",
			},
			{
				Title: "Saint valentin",
				Desc: "GPT-2 based valentines day letter generator.",
				Link: "https://github.com/jan25/saint-valentin",
			},
			{
				Title: "Portfolio",
				Desc:  "Minimal portfolio website using static HTML and Go templates. This website you're looking is the live demo :)",
				Link:  "https://github.com/jan25/jan25.github.io",
			},
		},
	},
	// About page content
	&AboutBlock{
		Meta: Meta{
			Name:    "about",
			Heading: "About",
		},
		Paragraphs: []string{
			`👋 Hi there! I'm Abhilash. I like Developer tooling and CI/CD spaces. Do checkout my <a href="/projects.html">Projects</a> section.`,

			`Currently and in the past, I enjoyed solving both product and infrastructure problems at large scale companies. I contribute to Open-source projects in spare time, mainly to Observability communities. I care about continuous learning and growth.`,

			`Besides work, I follow Formula 1, football and listen to ton of podcasts and radio. I also enjoy discovering cooking recipes, travelling to new places, long distance cycling and voluntary work near whereever i live.`,

			`Kia Ora! I found Kia Ora in this <a href="https://www.youtube.com/watch?v=zVDu0tJHTnY" target="_blank">video</a>.`,
		},
	},
}

// FooterLinks holds data for page footer
var FooterLinks = []FooterLink{
	{
		Title: "Github",
		URL:   "https://github.com/jan25",
	},
	{
		Title: "Linkedin",
		URL:   "https://www.linkedin.com/in/%F0%9F%91%8B-abhilash-gnan-0108b2105/",
	},
	{
		Title: "Twitter",
		URL:   "https://twitter.com/abhilashgnan",
	},
}

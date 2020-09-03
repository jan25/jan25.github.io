package main

// Blocks holds actual block data
var Blocks = []Block{
	&ProjectsBlock{
		Meta: Meta{
			Name:    "projects",
			Heading: "Projects",
		},
		Projects: []Project{
			{
				Title: "Termracer",
				Desc:  "Terminal based touch typing tutor written in Go. Works completely offline.",
				Link:  "https://github.com/jan25/termracer",
			},
			{
				Title: "Skripts",
				Desc:  "Execute node.js scripts inside a web browser. Comes with Node backend server and Vuejs frontend.",
				Link:  "https://github.com/jan25/skripts",
			},
			{
				Title: "Memory test",
				Desc:  "Test your short term memory skills using this simple React app.",
				Link:  "https://github.com/jan25/memory-test",
			},
			{
				Title: "Move my files",
				Desc:  "Simple CLI to organize files on your computer. Written in Python can be installed from pypi.",
				Link:  "https://github.com/jan25/move-my-files",
			},
			{
				Title: "Changelog gen",
				Desc:  "CLI tool implemented in Elixir to automate generation of CHANGELOG file for Github repositories.",
				Link:  "https://github.com/jan25/changeloggen",
			},
			{
				Title: "Game of life",
				Desc:  "React app showing Gosper's gilder gun, one of infinitely continuing patterns in Game of Life. R.I.P Conway.",
				Link:  "https://github.com/jan25/gameoflife",
			},
			{
				Title: "F1 Lights",
				Desc:  "Test your reaction times with this web app implemented using Vue.js.",
				Link:  "https://github.com/jan25/f1-lights",
			},
			{
				Title: "Sudoku",
				Desc:  "React app with fairly difficult puzzles that are pregenerated in Python.",
				Link:  "https://github.com/jan25/sandbox-2020/tree/master/sudoku",
			},
			{
				Title: "Track ISS",
				Desc:  "Close to real time position tracking of International space station. Implemented using Reactjs.",
				Link:  "https://github.com/jan25/track-iss",
			},
			{
				Title: "Valen10",
				Desc:  "Breakout game with emojis and silly gifs implemented on a Feb the 14th. Written in vanilla JS with support for offline mode.",
				Link:  "https://github.com/jan25/sandbox-2020/tree/master/valen10",
			},
			{
				Title: "HotRod",
				Desc:  "Jaeger's HotRod app rewritten in Python. Orignal demo app was written in Go.",
				Link:  "https://github.com/jan25/hotrod-python",
			},
		},
	},
	&AboutBlock{
		Meta: Meta{
			Name:    "about",
			Heading: "About",
		},
		Paragraphs: []string{
			`My name is Abhilash. I'm a Sofware engineer, mainly working on web applications. I've listed some of my projects in Projects section, do check it.`,

			`I work on all parts of web application development. Parts i enjoy most are at the intersection of product/application and infrastructure, such as CI/CD and Monitoring. Though currently focused on frontend and backend product development, my goal is to slowly transition into teams that build tooling to improve this intersection.`,

			`I contribute to opensource projects in spare time, mainly at observability communities, helping them with small-medium features, testing, documentation or anything i could do really in time i got. I enjoyed learning many things and not to mention giving back to community. I also got to interact with amazing engineers from different companies.`,

			`Besides technical work, I listen to tons of radio. I also enjoy cooking and discovering new recipes. I travel now and then to explore new places or simply exploring on my bike if not far. When I can i take part in voluntary work in my city.`,

			`Kia Ora! I found Kia Ora in this <a href="https://www.youtube.com/watch?v=zVDu0tJHTnY">video</a>. Have a great day :)`,
		},
	},
}

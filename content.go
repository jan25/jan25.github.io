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
				Title: "Termracer",
				Desc:  "Terminal based touch typing tutor written in Go. Works completely offline.",
				Link:  "https://github.com/jan25/termracer",
			},
			{
				Title: "Memory test",
				Desc:  "Test your short term memory skills using this simple React app.",
				Link:  "https://github.com/jan25/memory-test",
			},
			{
				Title: "Skripts",
				Desc:  "Execute node.js scripts inside a web browser. Comes with Node backend server and Vuejs frontend.",
				Link:  "https://github.com/jan25/skripts",
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
			`My name is Abhilash. I'm a Sofware engineer, mainly working on web applications. I've listed some of my projects in <a href="/projects.html">Projects</a> section, do check it :)`,

			`I work on all parts of web application development. Parts I enjoy most are at the intersection of product/application and infrastructure, such as CI/CD and Monitoring. Though currently focused on frontend and backend product development, my goal is to slowly transition into teams that build tooling to improve this intersection.`,

			`I contribute to Open-source projects in spare time, mainly to observability communities. I help them with small-medium features, testing, documentation or anything i can fit in my free time. I've enjoyed learning many things and not to mention giving back to community. I also got to interact with amazing engineers from different organizations.`,

			`Outside technical work, I follow F1 and listen to radio. I also enjoy cooking and discovering new recipes. I like traveling to explore new places or simply exploring neighborhoods on my bicycle if not far. When I can, I take part in voluntary work in my city.`,

			`Kia Ora! I found Kia Ora in this <a href="https://www.youtube.com/watch?v=zVDu0tJHTnY" target="_blank">video</a>. Have a great day!`,
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
		URL:   "https://www.linkedin.com/in/abhilash-gnan-0108b2105",
	},
	{
		Title: "Twitter",
		URL:   "https://twitter.com/abhilashgnan",
	},
}

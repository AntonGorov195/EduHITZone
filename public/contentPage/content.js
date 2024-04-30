// const isGuest = new URL(document.URL).searchParams.has("user", "guest");
//
// const userNameBox = document.getElementById("user-name-box");
// if (isGuest) {
// 	userNameBox.innerText = "שלום אורח!";
// } else {
//     console.error("Getting user's name is not implemented.");
// }

const coursesGroups = [
	{
		courseGroupName: "Popular",
		courses: [
			{
				name: "intro to C",
				thumbnail:
					"https://upload.wikimedia.org/wikipedia/commons/thumb/1/18/C_Programming_Language.svg/380px-C_Programming_Language.svg.png?20201031132917",
				href: "https://he.wikipedia.org/wiki/C_(%D7%A9%D7%A4%D7%AA_%D7%AA%D7%9B%D7%A0%D7%95%D7%AA)",
			},
		],
	},
];

// const courseGroupContainer = document.getElementById("course-group-container");
// 
// coursesGroups.forEach((courseGroup) => {
// 	_ = createAndAppend(courseGroupContainer, "h2", {
// 		textContent: courseGroup.courseGroupName,
// 	});
// 	const courseList = createAndAppend(courseGroupContainer, "ul", {
// 		classList: ["course-list"],
// 	});
//     courseGroup.courses.forEach((course) => {
//         const item = createAndAppend(courseList,)
//     });
// });
// 
// function createAndAppend() {}

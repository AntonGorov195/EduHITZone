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

/* <h2>מתמטיקה</h2>
<ul class="course-list">
	<li class="course-preview-item drop-box-shadow">
		<a class="course-preview"
			href="https://he.wikipedia.org/wiki/%D7%90%D7%99%D7%A0%D7%98%D7%92%D7%A8%D7%9C">
			<img class="course-preview-image"
				src="https://upload.wikimedia.org/wikipedia/commons/thumb/9/9f/Integral_example.svg/300px-Integral_example.svg.png">
			<div>
				אינפי 2
			</div>
		</a>
	</li>
</ul> */

const courseGroupContainer = document.getElementById("course-group-container");

coursesGroups.forEach((courseGroup) => {
	_ = createAndAppend(courseGroupContainer, "h2", {
		textContent: courseGroup.courseGroupName,
	});
	const courseList = createAndAppend(courseGroupContainer, "ul", {
		classList: ["course-list"],
	});
	courseGroup.courses.forEach((course) => {
		const item = createAndAppend(courseList, "li", {
			classList: ["course-preview-item drop-box-shadow"],
		});
		const a = createAndAppend(item, "a", {
			classList: "course-preview",
			href: course.href,
		});
		const img = createAndAppend(a, "img", {
			classList: "course-preview-image",
			onload: function () {
				img.classList.add("loaded");
			},
			src: course.thumbnail,
		});
		_ = createAndAppend(a, "div", { textContent: course.name });
	});
});
/**
 *
 * @param {HTMLElement} parent
 * @param {string} tagName
 * @param {any | undefined} props
 * @returns {HTMLElement}
 */
function createAndAppend(parent, tagName, props) {
	const newElem = document.createElement(tagName);

	if (props !== undefined) {
		for (const [key, value] of Object.entries(props)) {
			if (newElem[key] !== value) {
				newElem[key] = value;
			}
		}
	}
	parent.appendChild(newElem);
	return newElem;
}

let optionYearEles = document.querySelectorAll('.browser-default option');
for (let i = 0; i < optionYearEles.length; ++i) {
    optionYearEles[i].addEventListener('click', (e) => {
        console.log('Hello, world!');
    });
}
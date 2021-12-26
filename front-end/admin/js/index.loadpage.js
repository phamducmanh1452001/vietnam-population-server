$(document).ready(function () {
    let center_main_content = document.querySelector('#center-main-content');

    let cadresEle = document.getElementById('cadres');
    cadresEle.addEventListener('click', () => {
        // let center_main_content = document.querySelector('#center-main-content');
        // console.log(center_main_content);
        center_main_content.innerHTML = '';
        
        $(center_main_content).load('./cadres.html');
    });

    let citizenEle = document.getElementById('citizen');
    citizenEle.addEventListener('click', () => {
        // console.log(center_main_content);
        center_main_content.innerHTML = '';
        $(center_main_content).load('./citizen.html');
    });

    let personalInfoEle = document.querySelector('#personal-info');
    personalInfoEle.addEventListener('click', (e) => {
        center_main_content.innerHTML = '';
        $(center_main_content).load('./personal-info.html');
    })

    // let chartEle = document.getElementById('load-chart');
    // chartEle.addEventListener('click', () => {
    //     let center_main_content = document.querySelector('#center-main-content');
    //     // console.log(center_main_content);
    //     center_main_content.innerHTML = '';
    //     $(center_main_content).load('./chart.html');
    //     // console.log('Hello, chart')
    // });

    let chartYear = document.getElementById('chartYear');
    chartYear.addEventListener('click', () => {
        // let center_main_content = document.querySelector('#center-main-content');
        // console.log(center_main_content);
        center_main_content.innerHTML = '';
        $(center_main_content).load('/admin/chart/chartYear.html');
    });
    //chartSex
    let chartSex = document.getElementById('chartSex');
    chartSex.addEventListener('click', () => {
        // let center_main_content = document.querySelector('#center-main-content');
        // console.log(center_main_content);
        center_main_content.innerHTML = '';
        $(center_main_content).load('/admin/chart/chartSex.html');
    });
        //chartEthnic
    let chartEthnic = document.getElementById('chartEthnic');
    chartEthnic.addEventListener('click', () => {
        // let center_main_content = document.querySelector('#center-main-content');
        // console.log(center_main_content);
        center_main_content.innerHTML = '';
        $(center_main_content).load('/admin/chart/chartEthnic.html');
    });
    let chartCity = document.getElementById('chartCity');
    chartCity.addEventListener('click', () => {
        // let center_main_content = document.querySelector('#center-main-content');
        // console.log(center_main_content);
        center_main_content.innerHTML = '';
        $(center_main_content).load('/admin/chart/chartCity.html');
    });
    let chartAge = document.getElementById('chartAge');
    chartAge.addEventListener('click', () => {
        // let center_main_content = document.querySelector('#center-main-content');
        // console.log(center_main_content);
        center_main_content.innerHTML = '';
        $(center_main_content).load('/admin/chart/chartAge.html');
    });
});
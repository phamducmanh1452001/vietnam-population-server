$(document).ready(() => {

    let username = localStorage.getItem('user');
    let level = localStorage.getItem('level');

    fetch('https://www.phorifai.xyz/api/lower-cadres?page=1&limit=10', {
            method: 'GET',
            headers: {
                'Authorization': 'Bearer ' + localStorage.getItem('token')
            }
        })
        .then(response => response.json())
        .then(list_cadres => {
            // kh hiển thị chỉnh sửa quyền 
            if (parseInt(list_cadres['permission']) === 0) {
                document.querySelector('#btn-save-edit').style.display = 'none';
                document.querySelector('#all-permission').parentNode.style.display = 'none';
                document.querySelector('#all-disable-permission').parentNode.style.display = 'none';
            }
            console.log(list_cadres['permission'])
        })


    if (parseInt(level) === 0) {
        let provinceEle = document.querySelector('#province');
        let districtEle = document.querySelector('#district');
        let wardEle = document.querySelector('#ward');
        provinceEle.style.display = 'none';
        districtEle.style.display = 'none';
        wardEle.style.display = 'none';
        let addCitizeBtn = document.querySelector('#add-citizen');
        addCitizeBtn.style.display = 'none';

        document.querySelector('#page-number').innerHTML = '';
        showCitizenOfArea(localStorage.getItem('token'));
    }

    if (parseInt(level) === 1) {
        document.querySelector('#add-citizen').display = 'none';
        let provinceEle = document.querySelector('#province');
        provinceEle.style.display = 'none';
        let addCitizeBtn = document.querySelector('#add-citizen');
        addCitizeBtn.style.display = 'none';

        // district
        function showDistrict(province_id) { // data type 'province_id': string
            getDistricts(province_id)
                .then(data => {
                    let districts = data['data'];
                    document.querySelector('#district').innerHTML = '';
                    let districtEle = document.createElement('option');
                    districtEle.style.display = 'none';
                    districtEle.innerText = 'Chọn quận/huyện';
                    districtEle.value = '';
                    document.querySelector('#district').appendChild(districtEle);

                    document.querySelector('#ward').innerHTML = '';
                    let wardEle = document.createElement('option');
                    wardEle.style.display = 'none';
                    wardEle.innerText = 'Chọn phường/xã';
                    wardEle.value = '';
                    document.querySelector('#ward').appendChild(wardEle);

                    for (let j = 0; j < districts.length; ++j) {
                        let districtEle = document.createElement('option');
                        districtEle.innerText = districts[j]['name'];
                        districtEle.value = districts[j]['code'];
                        document.querySelector('#district').appendChild(districtEle);
                        // console.log(districtEle);

                    }
                    showCitizenOfArea(localStorage.getItem('token'));
                    // console.log(data);
                });
        }
        showDistrict(username); // them province_id tu tai khoan

        // ward
        function showWard(district_id) {
            getWards(district_id)
                .then(data => {
                    let wards = data['data'];
                    document.querySelector('#ward').innerHTML = '';
                    let wardEle = document.createElement('option');
                    wardEle.style.display = 'none';
                    wardEle.innerText = 'Chọn phường/xã';
                    wardEle.value = '';
                    document.querySelector('#ward').appendChild(wardEle);
                    for (let j = 0; j < wards.length; ++j) {
                        let wardEle = document.createElement('option');
                        wardEle.innerText = wards[j]['name'];
                        wardEle.value = wards[j]['code'];
                        document.querySelector('#ward').appendChild(wardEle);
                        // console.log(districtEle);
                    }
                    // // console.log(data);
                    $.post('https://www.phorifai.xyz/api/login', {
                        code: district_id,
                        password: district_id
                    }, (result) => {
                        // console.log(result.token);
                        document.querySelector('#page-number').innerHTML = '';
                        localStorage.setItem('token_1', result.token);
                        showCitizenOfArea(result.token);
                    });
                });
        }

        // handle event
        // function district_change() {
        //     let districtSelectedEle = document.querySelector('#district');
        //     showWard(districtSelectedEle.value);
        // }

        let DistrictSelect = document.querySelector('#district');
        DistrictSelect.onchange = (e) => {
            // console.log(e.target.value);
            showWard(e.target.value);
        }

        let WardSelect = document.querySelector('#ward');
        WardSelect.onchange = (e) => {
            $.post('https://www.phorifai.xyz/api/login', {
                code: e.target.value,
                password: e.target.value
            }, (result) => {
                // console.log(result.token);
                document.querySelector('#page-number').innerHTML = '';
                localStorage.setItem('token_2', result.token);
                showCitizenOfArea(result.token);
            });
        }
    }
    if (parseInt(level) === 2) {
        document.querySelector('#add-citizen').display = 'none';
        let districtEle = document.querySelector('#district');
        districtEle.style.display = 'none';
        let provinceEle = document.querySelector('#province');
        provinceEle.style.display = 'none';
        let addCitizeBtn = document.querySelector('#add-citizen');
        addCitizeBtn.style.display = 'none';

        // ward
        function showWard(district_id) {
            getWards(district_id)
                .then(data => {
                    let wards = data['data'];
                    document.querySelector('#ward').innerHTML = '';
                    let wardEle = document.createElement('option');
                    wardEle.style.display = 'none';
                    wardEle.innerText = 'Chọn phường/xã';
                    wardEle.value = '';
                    document.querySelector('#ward').appendChild(wardEle);
                    for (let j = 0; j < wards.length; ++j) {
                        let wardEle = document.createElement('option');
                        wardEle.innerText = wards[j]['name'];
                        wardEle.value = wards[j]['code'];
                        document.querySelector('#ward').appendChild(wardEle);
                        // console.log(districtEle);
                    }
                    // console.log(data);
                    showCitizenOfArea(localStorage.getItem('token'));

                    // $.post('https://www.phorifai.xyz/api/login', {
                    //     code: district_id,
                    //     password: district_id
                    // }, (result) => {
                    //     // console.log(result.token);
                    //     document.querySelector('#page-number').innerHTML = '';
                    //     localStorage.setItem('token_1', result.token);
                    //     // showCitizenOfArea(result.token);
                    // });
                });
        }
        showWard(username);

        let WardSelect = document.querySelector('#ward');
        WardSelect.onchange = (e) => {
            $.post('https://www.phorifai.xyz/api/login', {
                code: e.target.value,
                password: e.target.value
            }, (result) => {
                // console.log(result.token);
                document.querySelector('#page-number').innerHTML = '';
                localStorage.setItem('token_1', result.token);
                showCitizenOfArea(result.token);
            });
        }
    }
    if (parseInt(level) === 3) { // dành cho cán bộ xã
        let districtEle = document.querySelector('#district');
        districtEle.style.display = 'none';
        let provinceEle = document.querySelector('#province');
        provinceEle.style.display = 'none';
        let wardEle = document.querySelector('#ward');
        wardEle.style.display = 'none';

        showCitizenOfArea(localStorage.getItem('token'));
    }

})
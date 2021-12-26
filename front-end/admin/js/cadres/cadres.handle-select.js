$(document).ready(() => {

    let username = localStorage.getItem('user');
    let level = localStorage.getItem('level');

    // kh hiển thị chỉnh sửa quyền 
    if (parseInt(localStorage.getItem('permission')) === 0) {
        document.querySelector('#btn-save-edit').style.display = 'none';
        document.querySelector('#all-permission').parentNode.style.display = 'none';
        document.querySelector('#all-disable-permission').parentNode.style.display = 'none';

    }

    if (parseInt(level) === 0) {
        let provinceEle = document.querySelector('#province');
        let districtEle = document.querySelector('#district');
        let wardEle = document.querySelector('#ward');
        provinceEle.style.display = 'none';
        districtEle.style.display = 'none';
        wardEle.style.display = 'none';

        document.querySelector('#page-number').innerHTML = '';
        showCadresOfArea(localStorage.getItem('token'));
    }

    if (parseInt(level) === 1) {
        let provinceEle = document.querySelector('#province');
        let districtEle = document.querySelector('#district');
        let wardEle = document.querySelector('#ward');
        provinceEle.style.display = 'none';
        districtEle.style.display = 'none';
        wardEle.style.display = 'none';

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
                    showCadresOfArea(localStorage.getItem('token'));
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
                    // console.log(data);
                    $.post('https://www.phorifai.xyz/api/login', {
                        code: district_id,
                        password: district_id
                    }, (result) => {
                        // console.log(result.token);
                        document.querySelector('#page-number').innerHTML = '';
                        localStorage.setItem('token_1', result.token);
                        showCadresOfArea(result.token);
                    });
                });
        }

        // handle event
        // function district_change() {
        //     let districtSelectedEle = document.querySelector('#district');
        //     showWard(districtSelectedEle.value);
        // }

        let DistrictSelect = document.querySelector('#district');
        DistrictSelect.display = 'none';
        // DistrictSelect.onchange = (e) => {
        //     // console.log(e.target.value);
        //     showWard(e.target.value);
        // }
    }
    if (parseInt(level) === 2) {
        let districtEle = document.querySelector('#district');
        districtEle.style.display = 'none';
        let provinceEle = document.querySelector('#province');
        provinceEle.style.display = 'none';
        let wards = document.querySelector('#ward');
        wards.style.display = 'none';

        // ward
        function showWard(district_id) {
            fetch('https://evening-castle-31041.herokuapp.com/api/wards?district_code=' + district_id)
                .then(response => response.json())
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

                    $.post('https://www.phorifai.xyz/api/login', {
                        code: district_id,
                        password: district_id
                    }, (result) => {
                        // console.log(result.token);
                        document.querySelector('#page-number').innerHTML = '';
                        localStorage.setItem('token_1', result.token);
                    });
                });
        }
        showCadresOfArea(localStorage.getItem('token'));
        // showWard(username);
    }
    if (parseInt(level) === 3) { // dành cho cán bộ xã

    }

})
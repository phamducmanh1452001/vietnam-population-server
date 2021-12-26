// show province
function showProvince1() {
    getProvinces()
        .then((data) => {
            let provinces = data['data'];
            console.log(provinces[0]['name']);

            for (let i = 0; i < provinces.length; ++i) {
                let provinceEle = document.createElement('option');
                provinceEle.innerText = provinces[i]['name'];
                provinceEle.value = provinces[i]['code'];
                document.querySelector('#province1').appendChild(provinceEle);
            }
        });
}
showProvince1();

// district
function showDistrict1(province_id) { // data type 'province_id': string
    getDistricts(province_id)
        .then(data => {
            let districts = data['data'];
            document.querySelector('#district1').innerHTML = '';
            let districtEle = document.createElement('option');
            districtEle.style.display = 'none';
            districtEle.innerText = 'Chọn quận/huyện';
            districtEle.value = '';
            document.querySelector('#district1').appendChild(districtEle);

            document.querySelector('#ward1').innerHTML = '';
            let wardEle = document.createElement('option');
            wardEle.style.display = 'none';
            wardEle.innerText = 'Chọn phường/xã';
            wardEle.value = '';
            document.querySelector('#ward1').appendChild(wardEle);
            
            for (let j = 0; j < districts.length; ++j) {
                let districtEle = document.createElement('option');
                districtEle.innerText = districts[j]['name'];
                districtEle.value = districts[j]['code'];
                document.querySelector('#district1').appendChild(districtEle);
                // console.log(districtEle);
            }
            // console.log(data);
        });
}

// ward
function showWard1(district_id) {
    getWards(district_id)
        .then(data => {
            let wards = data['data'];
            document.querySelector('#ward1').innerHTML = '';
            let wardEle = document.createElement('option');
            wardEle.style.display = 'none';
            wardEle.innerText = 'Chọn phường/xã';
            wardEle.value = '';
            document.querySelector('#ward1').appendChild(wardEle);
            for (let j = 0; j < wards.length; ++j) {
                let wardEle = document.createElement('option');
                wardEle.innerText = wards[j]['name'];
                wardEle.value = wards[j]['code'];
                document.querySelector('#ward1').appendChild(wardEle);
                // console.log(districtEle);
            }
            // console.log(data);
        });
}

// handle event
function province_change_1() {
    let provinceSelectedEle = document.querySelector('#province1');
    // console.log(provinceSelectEle.value);
    showDistrict1(provinceSelectedEle.value);
}

// handle event
function district_change_1() {
    let districtSelectedEle = document.querySelector('#district1');
    showWard1(districtSelectedEle.value);
}
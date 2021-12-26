// handle login
let loginEle = document.querySelector('#btn-login');
loginEle.addEventListener('click', (e) => {
    localStorage.clear();

    let _data = {};
    _data.code = document.querySelector('#username').value.toString();
    _data.password = document.querySelector('#pwd').value.toString();

    $.ajax({
        method: 'POST',
        url: 'https://www.phorifai.xyz/api/login',
        data: _data,
        success: (response) => {
            // console.log(response.token); // return a string 
            if (response.token !== undefined) {
                localStorage.clear();
                localStorage.setItem('token', response.token);
                localStorage.setItem('user', _data.code);

                async function getProvinces(key_search) {
                    let response = await fetch('https://www.phorifai.xyz/api/provinces?page=1&limit=1000' + (key_search !== undefined ? `&key=${key_search}` : ``));
                    let data = await response.json();
                    return data;
                }

                async function getDistricts(province_code, key_search) {
                    let response = await fetch(`https://www.phorifai.xyz/api/districts?province_code=${province_code}&page=1&limit=1000` + (key_search !== undefined ? `&key=${key_search}` : ``));
                    let data = await response.json();
                    return data;
                }

                async function getWards(district_code, key_search) {
                    let response = await fetch(`https://www.phorifai.xyz/api/wards?district_code=${district_code}&page=1&limit=99999` + (key_search !== undefined ? `&key=${key_search}` : ``));
                    let data = await response.json();
                    return data;
                }

                if (_data.code === 'admin') {
                    localStorage.setItem('name_area', 'Việt Nam');
                    localStorage.setItem('level', '0')
                    localStorage.setItem('permission', response.permission);
                    window.location = '../../admin/index.html';
                }

                getProvinces(_data.code)
                    .then(data => {
                        let provinces = data['data'][0];
                        localStorage.setItem('name_area', provinces['name']);
                        localStorage.setItem('level', '1')
                        localStorage.setItem('permission', response.permission);
                        window.location = '../../admin/index.html';
                    })


                getProvinces()
                    .then(data => {
                        let provinces = data['data'];
                        for (let i = 0; i < provinces.length; ++i) {
                            getDistricts(provinces[i]['code'], _data.code)
                                .then(data1 => {
                                    let districts = data1['data'][0];
                                    localStorage.setItem('name_area', `${districts['name']} - ${provinces[i]['name']}`);
                                    localStorage.setItem('level', '2');
                                    localStorage.setItem('permission', response.permission);
                                    window.location = '../../admin/index.html';
                                })
                        }
                    })

                getProvinces()
                    .then(data => {
                        let provinces = data['data'];
                        for (let i = 0; i < provinces.length; ++i) {
                            getDistricts(provinces[i]['code'])
                                .then(data1 => {
                                    let districts = data1['data'];
                                    for (let j = 0; j < districts.length; ++j) {
                                        getWards(districts[j]['code'], _data.code)
                                            .then(data2 => {
                                                let wards = data2['data'][0];
                                                localStorage.setItem('name_area', `${wards['name']} - ${districts[j]['name']} - ${provinces[i]['name']}`);
                                                localStorage.setItem('level', '3');
                                                localStorage.setItem('permission', response.permission);
                                                window.location = '../../admin/index.html';
                                            })
                                    }
                                })
                        }
                    })

                // if (_data.code.length == 2) {
                //     getProvinces(_data.code)
                //         .then(data => {
                //             if (data['data'].length > 0) {
                //                 let provinces = data['data'][0];
                //                 localStorage.setItem('name_area', provinces['name']);
                //                 window.location = '../../admin/index.html';
                //             }
                //         })
                // }

                // if (_data.code.length === 3) {
                //     getProvinces()
                //         .then(data => {
                //             let provinces = data['data'];
                //             // console.log(provinces);
                //             for (let i = 0; i < provinces.length; ++i) {
                //                 getDistricts(provinces[i]['code'], _data.code)
                //                     .then(data1 => {
                //                         if (data1['data'].length > 0) {
                //                             let districts = data1['data'][0];
                //                             localStorage.setItem('name_area', `${districts['name']} - ${provinces[i]['name']}`);
                //                             window.location = '../../admin/index.html';
                //                         }
                //                     })
                //             }
                //         })
                // }
                // if (_data.code.length === 5) {
                //     getProvinces()
                //     .then(data => {
                //         let provinces = data['data'];
                //         for (let i = 0; i < provinces.length; ++i) {
                //             getDistricts(provinces[i]['code'])
                //             .then(data1 => {
                //                 let districts = data1['data'];
                //                 // console.log(districts);
                //                 for (let j = 0; j < districts.length; ++j) {
                //                     getWards(districts[j]['code'], _data.code)
                //                     .then(data2 => {
                //                         // if (data2['data'].length > 0) {
                //                         //     let wards = data2['data'][0];
                //                         //     localStorage.setItem('name_area', `${wards['name']} - ${districts[j]['name']} - ${provinces[i]['name']}`);
                //                         //     window.location = '../../admin/index.html';
                //                         // }
                //                         console.log(data2);
                //                     })
                //                 }
                //             })
                //         }
                //     })
                // }
            } else {
                alert('Đăng nhập thất bại!');
            }

        },
        error: (er) => {
            alert('Đăng nhập thất bại!');
        }
    })

});
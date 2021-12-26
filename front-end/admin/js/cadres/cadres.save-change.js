$(document).ready(() => {

    function updatePermission(_code, _permission) {
        return $.ajax({
            method: 'POST',
            url: 'https://www.phorifai.xyz/api/change-cadre-permission',
            headers: {
                'Authorization': 'Bearer ' + localStorage.getItem('token')
            },
            data: {
                code: _code.toString(),
                permission: parseInt(_permission)
            }
        });
    }

    async function getCodeCadres(page, limit) {
        let response = await fetch(`https://www.phorifai.xyz/api/lower-cadres?page=${page}&limit=${limit}`, {
            method: 'GET',
            headers: {
                'Authorization': 'Bearer ' + localStorage.getItem('token')
            }
        });
        let data = await response.json();
        return data;
    }

    async function returnMessage(_code, _permission) {
        let response = await updatePermission(_code, _permission);
        // console.log(response);
        return response;
    }

    let allPermissionCheckbox = document.querySelector('#all-permission');
    allPermissionCheckbox.onclick = (e) => {
        if (e.target.checked == true) {
            let check = false;
            // console.log(e.target.value);
            getCodeCadres(1, 99999)
                .then(data => {
                    // console.log(data);
                    for (let i = 0; i < data['data'].length; ++i) {
                        returnMessage(data['data'][i]['code'], 1);
                        check = true;
                    }
                    if (check) {
                        e.target.checked = false;
                        document.querySelector('#page-number').innerHTML = '';
                        showCadresOfArea(localStorage.getItem('token'));
                        alert('Thay đổi thành công!');
                    }
                })

        }
    }

    let allDisablePermissionCheckbox = document.querySelector('#all-disable-permission');
    allDisablePermissionCheckbox.onclick = (e) => {
        if (e.target.checked == true) {
            let check = false;
            // console.log(e.target.value);
            getCodeCadres(1, 99999)
                .then(data => {
                    // console.log(data);
                    for (let i = 0; i < data['data'].length; ++i) {
                        returnMessage(data['data'][i]['code'], 0);
                        check = true;
                    }
                    if (check) {
                        e.target.checked = false;
                        document.querySelector('#page-number').innerHTML = '';
                        showCadresOfArea(localStorage.getItem('token'));

                        alert('Thay đổi thành công!');
                    }
                })

        }
    }

    let saveChangeButton = document.querySelector('#btn-save-edit');
    let check = false;
    saveChangeButton.addEventListener('click', (e) => {
        // console.log(map);
        map.forEach((value, key) => {
            returnMessage(key, value);
            check = true;
        });
        if (check) {
            alert('Thay đổi thành công!');
        }
    });
})
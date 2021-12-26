$(document).ready(() => {
    let religions = {'0': 'Không', '1': 'Phật giáo', '2': 'Công giáo', '3': 'Hồi giáo', '4': 'Tin lành', '5': 'Cao Đài', '6': 'Tôn giáo dân gian', '7': 'Hoà hảo', '8': 'Khác'};
    let searchInput = document.querySelector('#search');
    let searchBtn = document.querySelector('#btn-search');

    function getUrl(page, limit, avatar) {
        return `https://www.phorifai.xyz/api/citizens?page=${page.toString()}&limit=${limit.toString()}`;
    }

    async function getCitizen(_token, key_search) { // modify
        let response = await fetch(`https://www.phorifai.xyz/api/citizens?page=1&limit=99999&key=${key_search}`, {
            method: 'GET',
            headers: {
                'Authorization': 'Bearer ' + _token
            }
        });

        let data = await response.json();
        return data;
    }

    searchBtn.onclick = () => {
        let showDiv = document.querySelector('#show-info-detail');
        showDiv.className = 'card shadow mb-1';
        let subShowDiv = showDiv.querySelectorAll('div')[0];
        subShowDiv.innerHTML = '';

        let valueInput = searchInput.value.toString();

        getCitizen(localStorage.getItem('token'), valueInput)
            .then(data => {
                let citizen = data['data'][0];
                let spans = new Array(10);
                for (let i = 0; i < spans.length; ++i) {
                    spans[i] = document.createElement('span');
                    spans[i].style.display = 'block';
                    spans[i].style.padding = '3px 5px';
                }
                // spans[0].innerHTML = `Code: ${citizen['code']}`;
                // spans[1].innerHTML = `Họ tên: ${citizen['name']}`;
                // spans[2].innerHTML = `Tuổi: ${cadres['age']}`;
                // spans[3].innerHTML = `Giới tính: ${citizen['gender']}`;
                // spans[4].innerHTML = `Địa chỉ:`;
                // spans[5].innerHTML = `Tôn giáo: ${citizen['religion']}`;
                // for (let i = 0; i < spans.length; ++i) {
                //     subShowDiv.appendChild(spans[i]);
                //     // console.log(spans[i])
                // }
                spans[0].innerHTML = `Mã số: ${citizen['code']}`;
                spans[1].innerHTML = `Họ tên: ${citizen['first_name']} ${citizen['middle_name']} ${citizen['last_name']}`;
                spans[2].innerHTML = `Tuổi: ${citizen['age']}`;
                spans[3].innerHTML = `Giới tính: ` + (citizen['gender'] === 'M' ? 'Nam' : 'Nữ');
                spans[4].innerHTML = `Địa chỉ: ${citizen['address']}`;
                spans[5].innerHTML = `Tôn giáo: ${religions[citizen['religion_id']]}`;
                if (citizen['collaborator_name'] !== '' && citizen['collaborator_phone'] !== '') {
                    spans[6].innerHTML = `Tên cộng tác viên: ${citizen['collaborator_name']}`;
                    spans[7].innerHTML = `Số điện thoại: ${citizen['collaborator_phone']}`;
                }
                spans[8].innerHTML = `Địa chỉ tạm trú: ${citizen['temporary_address']}`;
                spans[9].innerHTML = `Ảnh`;
                for (let i = 0; i < spans.length; ++i) {
                    subShowDiv.appendChild(spans[i]);
                }

            })
    }
})
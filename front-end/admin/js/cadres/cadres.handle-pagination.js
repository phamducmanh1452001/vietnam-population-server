function showCadresOfArea(_token, key_search) {
    map = new Map();
    // show personal detail
    function showPersonalDetail(cadres, permis) {
        let showDiv = document.querySelector('#show-info-detail');
        showDiv.className = 'card shadow mb-1';
        let subShowDiv = showDiv.querySelectorAll('div')[0];
        subShowDiv.innerHTML = '';
        let spans = new Array(7);
        for (let i = 0; i < spans.length; ++i) {
            spans[i] = document.createElement('span');
            spans[i].style.display = 'block';
            spans[i].style.padding = '3px 5px';
        }
        spans[0].innerHTML = `Mã khu vực: ${cadres['code']}`;
        spans[1].innerHTML = `Chức vụ: ${cadres['name']}`;
        spans[2].innerHTML = `Tuổi: ${cadres['age']}`;
        spans[3].innerHTML = `Số điện thoại: ${cadres['phone']}`;
        spans[4].innerHTML = `Email: ${cadres['email']}`;
        spans[5].innerHTML = 'Quyền chỉnh sửa: ' + (parseInt(permis) === 1 ? 'Có' : 'Không');
        spans[6].innerHTML = `Khu vực quản lý: ${cadres['subdivision']['name']}`;
        for (let i = 0; i < spans.length; ++i) {
            subShowDiv.appendChild(spans[i]);
        }
    }

    // set permission 
    function addEventEdit(td) {
        td.onclick = () => {
            let permission = td.parentNode.querySelector('td:nth-child(5)');
            // document.querySelectorAll('#permission option')[parseInt(permission)].selected = true;
            if (permission.innerHTML === '1') {
                permission.innerHTML = '0';
                let _code = td.parentNode.querySelectorAll('td')[0].innerHTML.toString();
                let _permission = td.parentNode.querySelector('td:nth-child(5)').innerHTML.toString();
                map.set(_code, _permission);
            } else {
                permission.innerHTML = '1';
                let _code = td.parentNode.querySelectorAll('td')[0].innerHTML.toString();
                let _permission = td.parentNode.querySelector('td:nth-child(5)').innerHTML.toString();
                map.set(_code, _permission);
            }
            // console.log(map);

            // let savePermissionButton = document.querySelector('#save-permission');
            // savePermissionButton.addEventListener('click', (e) => {
            //     td.parentNode.querySelector('td:nth-child(5)').innerHTML = document.querySelector('#permission').value;
            //     // console.log(map);
            // })
        }
    }

    // show personal information
    function addEventPersonalDetail(tr, cadres) {
        tr.querySelector('td:nth-child(2)').onclick = () => {
            showPersonalDetail(cadres, tr.querySelector('td:nth-child(5)').innerHTML);
        }
    }

    function showCadres(list_cadres) { // list_cadres is an array consists of objects
        document.querySelector('tbody').innerHTML = '';
        for (let i = 0; i < list_cadres['data'].length; ++i) {
            let tr = document.createElement('tr');
            let td = new Array(6);
            for (let j = 0; j < td.length; ++j) {
                td[j] = document.createElement('td');
            }
            // console.log(list_cadres['data'][i]['name']);
            td[0].innerHTML = list_cadres['data'][i]['code'];
            td[1].innerHTML = list_cadres['data'][i]['name'];
            td[2].innerHTML = list_cadres['data'][i]['age'];
            td[3].innerHTML = list_cadres['data'][i]['subdivision']['name'];
            td[4].innerHTML = list_cadres['data'][i]['permission'];
            td[5].className = '';
            let span = document.createElement('span');
            // span.setAttribute('data-toggle', 'modal');
            // span.setAttribute('data-target', '#myModal');
            span.innerHTML = 'Edit <i class="far fw fa-edit"></i>';
            td[5].appendChild(span);
            addEventEdit(td[5]);
            for (let j = 0; j < td.length; ++j) {
                tr.appendChild(td[j]);
            }
            addEventPersonalDetail(tr, list_cadres['data'][i])
            document.querySelector('tbody').appendChild(tr);
        }
    }

    function countPage(numofrecords, numofdisplays) {
        if (parseInt(numofdisplays) > parseInt(numofrecords)) {
            return 1;
        }
        if (parseInt(numofrecords) % parseInt(numofdisplays) == 0) {
            return parseInt(numofrecords) / parseInt(numofdisplays);
        }
        return parseInt(parseFloat(numofrecords) / parseFloat(numofdisplays)) + 1;
    }

    fetch('https://www.phorifai.xyz/api/lower-cadres?page=1&limit=10', {
            method: 'GET',
            headers: {
                'Authorization': 'Bearer ' + _token
            }
        })
        .then(response => response.json())
        .then(list_cadres => {

            // tổng khu vực
            document.querySelector('#total_area').innerHTML = `Tổng: ${list_cadres['amount']} khu vực`;

            // console.log(list_cadres);
            let cntPages = countPage(list_cadres['amount'], 10);
            // console.log(cntPages);
            let option = new Array(cntPages);
            for (let i = 0; i < cntPages; ++i) {
                option[i] = document.createElement('option');
                option[i].value = (i + 1).toString();
                option[i].innerHTML = (i + 1).toString();
                document.querySelector('#page-number').appendChild(option[i]);
            }
            showCadres(list_cadres);
        })

    let numOfPagesEle = document.querySelector('#page-number');
    let numOfDisplaysEle = document.querySelector('#data_tables');

    numOfPagesEle.onchange = (e) => {
        let optionEle = e.target;
        fetch('https://www.phorifai.xyz/api/lower-cadres?page=' + optionEle.value.toString() + '&limit=' + numOfDisplaysEle.value.toString(), {
                method: 'GET',
                headers: {
                    'Authorization': 'Bearer ' + _token
                }
            })
            .then(response => response.json())
            .then(list_cadres => {
                showCadres(list_cadres);
            })
    }

    numOfDisplaysEle.onchange = (e) => {
        let optionEle = e.target;
        // console.log(optionEle.value);

        fetch('https://www.phorifai.xyz/api/lower-cadres?page=' + numOfPagesEle.value.toString() + '&limit=' + optionEle.value.toString(), {
                method: 'GET',
                headers: {
                    'Authorization': 'Bearer ' + _token
                }
            })
            .then(response => response.json())
            .then(list_cadres => {
                document.querySelector('#page-number').innerHTML = '';
                let cntPages = countPage(list_cadres['amount'], optionEle.value);
                let option = new Array(cntPages);
                for (let i = 0; i < cntPages; ++i) {
                    option[i] = document.createElement('option');
                    option[i].value = (i + 1).toString();
                    option[i].innerHTML = (i + 1).toString();
                    document.querySelector('#page-number').appendChild(option[i]);
                }
                fetch('https://www.phorifai.xyz/api/lower-cadres?page=' + numOfPagesEle.value.toString() + '&limit=' + optionEle.value.toString(), {
                        method: 'GET',
                        headers: {
                            'Authorization': 'Bearer ' + _token
                        }
                    })
                    .then(response1 => response1.json())
                    .then(list_cadres1 => {
                        document.querySelector('#page-number').innerHTML = '';
                        let cntPages1 = countPage(list_cadres1['amount'], optionEle.value);
                        let option1 = new Array(cntPages1);
                        for (let i = 0; i < cntPages1; ++i) {
                            option1[i] = document.createElement('option');
                            option1[i].value = (i + 1).toString();
                            option1[i].innerHTML = (i + 1).toString();
                            document.querySelector('#page-number').appendChild(option1[i]);
                        }
                        showCadres(list_cadres1);
                    })
            })
    }
}
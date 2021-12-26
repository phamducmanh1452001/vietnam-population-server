$(document).ready(() => {
    if (parseInt(localStorage.getItem('permission')) === 0) {
        document.querySelector('#add-citizen').style.display = 'none';
    }
    document.querySelector('#info-collaborator').className = 'd-none';
    document.querySelector('#descrip-religion').className = 'd-none';

    let collaboratorSelect = document.querySelector('#collaborator');
    collaboratorSelect.onchange = (e) => {
        if (e.target.value === '1') {
            document.querySelector('#info-collaborator').className = 'd-flex justify-content-between';
        } else {
            document.querySelector('#info-collaborator').className = 'd-none';
        }
    }

    let religionSelect = document.querySelector('#religions');
    religionSelect.onchange = (e) => {
        if (e.target.value === '8') {
            document.querySelector('#descrip-religion').className = 'form-group';
        } else {
            document.querySelector('#descrip-religion').className = 'd-none';
        }
    }

    async function checkDuplicateCitizen(page, limit, key_search) {
        let response = await fetch(`https://www.phorifai.xyz/api/citizens?page=${page}&limit=${limit}&key=${key_search}`, {
            method: 'GET',
            headers: {
                'Authorization': 'Bearer ' + localStorage.getItem('token')
            }
        });
        let data = await response.json();
        return data;
    }

    function currentDay() {
        let today = new Date();
        let dd = String(today.getDate()).padStart(2, '0');
        let mm = String(today.getMonth() + 1).padStart(2, '0'); //January is 0!
        let yyyy = today.getFullYear();

        today = `${yyyy}-${mm}-${dd}`;
        return today;
    }

    function getAge(dateString) {
        let today = new Date();
        let birthDate = new Date(dateString);
        let age = today.getFullYear() - birthDate.getFullYear();
        let m = today.getMonth() - birthDate.getMonth();
        if (m < 0 || (m === 0 && today.getDate() < birthDate.getDate())) {
            age--;
        }
        return age;
    }

    let addCitizenBtn = document.querySelector('#btn-add-citizen');
    addCitizenBtn.onclick = (e) => {
        e.preventDefault();
        let first_name = document.querySelector('#first_name').value.trim().replace('  ', ' ');
        let middle_name = document.querySelector('#middle_name').value.trim().replace('  ', ' ');
        let last_name = document.querySelector('#last_name').value.trim().replace('  ', ' ');
        let gender = (document.querySelector('#gender').value === 'male' ? 'M' : 'F');
        let dob = document.querySelector('#dob').value;
        let province = document.querySelector('#province1').options[document.querySelector('#province1').selectedIndex].text;
        let district = document.querySelector('#district1').options[document.querySelector('#district1').selectedIndex].text;
        let ward = document.querySelector('#ward1').options[document.querySelector('#ward1').selectedIndex].text;
        let weight = document.querySelector('#weight').value;
        let religion_id = document.querySelector('#religions').value;
        let major = document.querySelector('#major').value;
        let isCollaborator = document.querySelector('#collaborator').value;
        let collaborator_name = document.querySelector('#collaborator_name').value.trim().replace('  ', ' ');
        let collaborator_phone = document.querySelector('#collaborator_phone').value.trim();
        let temporary_address = `${ward} - ${district} - ${province}`;
        let code = document.querySelector('#cccd').value.trim();
        let file_name = document.querySelector('#file-portrait').value;

        let formDetail = [];
        formDetail.push(first_name, middle_name, last_name, gender, dob, weight, religion_id, major, isCollaborator, collaborator_name, collaborator_phone, temporary_address, code, file_name)
        // console.log(checkFormEmpty(formDetail));

        console.log(formData);

        if ((code.length !== 8 && code.length !== 12) || checkFormEmpty(formDetail)) {
            alert('Thông tin không hợp lệ');
        } else {
            let personal = {};
            personal.code = code;
            personal.first_name = first_name;
            personal.middle_name = middle_name;
            personal.last_name = last_name;
            personal.gender = gender;
            personal.date_of_birth = dob;
            personal.date_of_joining = currentDay();
            personal.religion_id = religion_id;
            personal.collaborator_name = collaborator_name;
            personal.collaborator_phone = collaborator_phone;
            personal.major = major;
            personal.temporary_address = temporary_address;

            checkDuplicateCitizen(1, 99999, code)
                .then(data => {
                    // console.log(data);
                    // if (data['data'].length > 0) {
                    //     alert('Mã số đã tồn tại!');
                    // } else {
                    // $.post('https://www.phorifai.xyz/api/upload', 
                    // {
                    //     headers: {
                    //         'Authorization': 'Bearer ' + localStorage.getItem('token'),
                    //         'Content-Type': 'multipart/form-data'
                    //     },
                    //     body: formData
                    // },
                    // (img => {
                    //     console.log(img)
                    // })


                    // )
                    fetch('https://www.phorifai.xyz/api/upload', {
                            mode: 'no-cors',
                            method: 'POST',
                            headers: {
                                'Authorization': 'Bearer ' + localStorage.getItem('token'),
                            },
                            body: formData
                        })
                        .then(response => response.json())
                        .then(img => {
                            console.log(img['name'].split('/')[1].toString());
                            personal.avatar = img['name'].split('/')[1].toString();
                            // fetch('https://www.phorifai.xyz/api/add-citizen', {
                            //         method: 'POST',
                            //         headers: {
                            //             'Authorization': 'Bearer ' + localStorage.getItem('token')
                            //         },
                            //         body: JSON.stringify(personal)
                            //     })
                            //     .then(response1 => response1.json())
                            //     .then(data => {
                            //         // console.log(data);
                            //         alert('Thêm thành công!');
                            //     })
                        })


                    // $.post({
                    //     method: 'POST',
                    //     url: 'https://www.phorifai.xyz/api/upload',
                    //     headers: {
                    //         'Authorization': 'Bearer ' + localStorage.getItem('token'),
                    //         'Content-Type': 'multipart/form-data'
                    //     },
                    //     data: formData,
                    //     success: (response) => {
                    //         console.log(response);
                    //     }
                    // })

                    // $.ajax({
                    //     method: 'POST',
                    //     url: 'https://www.phorifai.xyz/api/upload',
                    //     headers: {
                    //         'Authorization': 'Bearer ' + localStorage.getItem('token'),
                    //         'Content-Type': 'multipart/form-data'
                    //     },
                    //     data: formData,
                    //     success: (response) => {
                    //         console.log(response)
                    //         // personal.avatar = 
                    //     }
                    // })

                    $.post({
                        mode: 'no-cors',
                        method: 'POST',
                        url: 'https://www.phorifai.xyz/api/add-citizen',
                        headers: {
                            'Authorization': 'Bearer ' + localStorage.getItem('token')
                        },
                        data: personal,
                        success: (response => {
                            // console.log(personal);
                            // console.log('Hello, world')
                            alert('Thêm thành công!');
                        })
                    })



                    // }
                    // console.log(data);
                    // console.log(personal);
                })


        }
    }


})
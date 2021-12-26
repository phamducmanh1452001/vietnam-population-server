function checkFormEmpty(formDetail) { // formDetail is an array
    // let arr = ['first_name', 'middle_name', 'last_name', 'dob', 'province_code', 'district_code', 'ward_code', 'weight', 'major', 'collaborator_name', 'collaborator_phone', 'temporary_address', 'code', 'img'];
    // if (first_name != '' && middle_name != )

    let len = formDetail.length;
    let ok = false;
    let ok1 = false;

    if (formDetail[11] === '0') {
        for (let i = 0; i < len; ++i) {
            if (formDetail[i] === '' && i !== 12 && i != 13) {
                ok = true;
            }
        }
    }
    if (formDetail[11] === '1') {
        for (let i = 0; i < len; ++i) {
            if (formDetail[i] === '') {
                ok1 = true;
            }
        }
    }
    if (ok || ok1) {
        return true;
    }
    if (!ok && !ok1) {
        return false;
    }
}
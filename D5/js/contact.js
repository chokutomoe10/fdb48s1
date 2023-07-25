function submitData() {

    let name = document.getElementById("input-name").value
    let email = document.getElementById("input-email").value
    let phone = document.getElementById("input-phone").value
    let subject = document.getElementById("input-subject").value
    let message = document.getElementById("input-message").value

    let mailerData = {
        name: name,
        email,
        phone,
        subject,
        message
    }

    console.log(mailerData)

    if (name === "") {
        return alert('Name harus diisi!')
    } else if (email === "") {
        return alert('Email harus diisi!')
    } else if (phone === "") {
        return alert('Phone harus diisi!')
    } else if (subject === "") {
        return alert('Subject harus diisi!')
    } else if (message === "") {
        return alert('Message harus diisi!')
    }

    let a = document.createElement('a')
    a.href = `mailto:${email}?subject=${subject}&body=Halo, nama saya ${name},\n${message}, silahkan menghubungi saya di nomor berikut : ${phone}`
    a.click()
}
class Testimonial {
    #quote = ""
    #image = ""

    constructor(quote, image) {
        this.#quote = quote
        this.#image = image
    }

    get quote() {
        return this.#quote
    }

    get image() {
        return this.#image
    }

    get user() {
        throw new Error('there is must be user to make testimonials')
    }

    get testimonialHTML() {
        return `<div class="testimonial">
            <img src="${this.image}" />
            <p class="quote">"${this.quote}"</p>
            <p class="author">- ${this.user}</p>
        </div>
        `
    }
}

class UserTestimonial extends Testimonial {
    #user = ""

    constructor(user, quote, image) {
        super(quote, image)
        this.#user = user
    }

    get user() {
        return this.#user
    }
}

class CompanyTestimonial extends Testimonial {
    #company = ""

    constructor(company, quote, image) {
        super(quote, image)
        this.#company = company
    }

    get user() {
        return this.#company + " company"
    }
}

const testimonial1 = new UserTestimonial("Zenin", "Wah, luar biasa", "https://media.discordapp.net/attachments/305431915178098688/614225458719883298/Capture.PNG")
const testimonial2 = new UserTestimonial("Juro", "Asli keren banget", "https://data.1freewallpapers.com/download/devil-may-cry-dante-vs-vergil.jpg")
const testimonial3 = new CompanyTestimonial("Fang", "Menarik", "https://www.devilmaycry.com/5se/assets/images/topics-panel-vergil.jpg")

let testimonialData = [testimonial1, testimonial2, testimonial3]

let testimonialHTML = ""

for (let i = 0; i < testimonialData.length; i++) {
    testimonialHTML += testimonialData[i].testimonialHTML
}

document.getElementById("testimonials").innerHTML = testimonialHTML
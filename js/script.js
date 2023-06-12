"use strict"

const getDogImage = async (URLAddress) => {
  try {
    const response = await fetch(URLAddress)
    const json = await response.json()
    document.getElementById(
      "dog"
    ).innerHTML = `<img src="${json.message}" alt="Dog Image">`
  } catch (err) {
    console.log(err)
  }
}

getDogImage("https://dog.ceo/api/breeds/image/random")

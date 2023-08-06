const { createCommand } = require("commander");
const axios = require("axios");
const FormData = require("form-data");
const colors = require("colors");

function shorten(link, options) {
  let formData = new FormData();
  formData.append("link", link);

  if (options.password) formData.append("password", options.password);

  if (!options.instance) options.instance = "https://s.oriondev.fr";
  if (!options.instance.startsWith("http")) options.instance = `https://${options.instance}`;

  axios.get(`${options.instance}/api/quectoCheck`).then((res) => {
    if (res.data.data.quecto) {
      axios.post(`${options.instance}/api/shorten`, formData, {}).then((res) => {
        console.log(`Shortened link: ${res.data.data.shorten}`.brightGreen);
      })
    } else {
      console.log("This instance is not a Quecto instance.".brightRed);
    }
  }).catch((e) => {
    console.log("This instance is not a Quecto instance.".brightRed);
  })
}

module.exports = createCommand("shorten")
  .description("Shorten a URL")
  .argument("<link>", "The link to shorten")
  .option("-i, --instance <instance>", "The instance to use")
  .option("-p, --password <password>", "The password for shortened link")
  .action(shorten)
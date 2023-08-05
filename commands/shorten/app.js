const { createCommand } = require("commander");
const axios = require("axios");
const FormData = require("form-data");
const colors = require("colors");

function shorten(link, instance) {
  let formData = new FormData();
  formData.append("link", link);

  if (!instance.instance) instance.instance = "https://s.oriondev.fr";
  if (!instance.instance.startsWith("http")) instance.instance = `https://${instance.instance}`;

  axios.get(`${instance.instance}/api/quectoCheck`).then((res) => {
    if (res.data.data.quecto) {
      axios.post(`${instance.instance}/api/shorten`, formData, {}).then((res) => {
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
  .action(shorten)
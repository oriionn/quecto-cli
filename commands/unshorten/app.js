const {createCommand} = require("commander");
const axios = require("axios");
const colors = require("colors");

function unshorten(link, instance) {

  if (!instance.instance) instance.instance = "https://s.oriondev.fr";
  if (!instance.instance.startsWith("http")) instance.instance = `https://${instance.instance}`;

  axios.get(`${instance.instance}/api/quectoCheck`).then((res) => {
    if (res.data.data.quecto) {
      axios.get(`${instance.instance}/api/s/${link.split("/")[link.split("/").length - 1]}`, {}).then((res) => {
        console.log(`Unshortened link: ${res.data.data.original}`.brightGreen);
      }).catch((e) => {
        console.log("This link is not a from this instance.".brightRed);
      })
    } else {
      console.log("This instance is not a Quecto instance.".brightRed);
    }
  }).catch((e) => {
    console.log("This instance is not a Quecto instance.".brightRed);
  })
}

module.exports = createCommand("unshorten")
    .description("Unshorten a URL")
    .argument("<link>", "The link to unshorten")
    .option("-i, --instance <instance>", "The instance to use")
    .action(unshorten)
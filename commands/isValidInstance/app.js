const {createCommand} = require("commander");
const axios = require("axios");
const colors = require("colors");

function verifyInstance(instance) {
  if (!instance.startsWith("http")) instance.instance = `https://${instance.instance}`;

  axios.get(`${instance}/api/quectoCheck`).then((res) => {
    if (res.data.data.quecto) {
      console.log("This instance is a Quecto instance.".brightGreen);
    } else {
      console.log("This instance is not a Quecto instance.".brightRed);
    }
  }).catch((e) => {
    console.log("This instance is not a Quecto instance.".brightRed);
  })
}

module.exports = createCommand("ivi")
    .description("Verify if a domain is a Quecto instance")
    .argument("<instance>", "The domain of the instance")
    .action(verifyInstance)
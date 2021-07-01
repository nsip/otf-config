// includes some variables from caller 'template:'

import { getForm_NatsStreaming } from "./natsstreaming.js";
import { getForm_Nias3 } from "./nias3.js";
import { getForm_Benthos } from "./benthos.js";
import { getForm_Reader } from "./reader.js";
import { getForm_Align } from "./align.js";
import { getForm_TxtClassifier } from "./txtclassifier.js";
import { getForm_Level } from "./level.js";
import { getForm_Weight } from "./weight.js";
import { getForm_Hub } from "./hub.js";

export function getForm() {

  let form_head = `<div id="mainform">
  <h4 style="text-align: center;">{{title}}</h4>`;

  let form_tail = `<br></div>`;

  return (
    form_head +
    getForm_NatsStreaming() +
    getForm_Nias3() +
    getForm_Benthos() +
    getForm_Reader() +
    getForm_Align() +
    getForm_TxtClassifier() +
    getForm_Level() +
    getForm_Weight() +
    getForm_Hub() +
    form_tail
  );
}

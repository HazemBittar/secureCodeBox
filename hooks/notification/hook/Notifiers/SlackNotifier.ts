// SPDX-FileCopyrightText: the secureCodeBox authors
//
// SPDX-License-Identifier: Apache-2.0

import { NotifierType } from "../NotifierType";
import { AbstractWebHookNotifier } from "./AbstractWebHookNotifier";
import { Finding } from "../model/Finding";
import axios from "axios";
import { NotificationChannel } from "../model/NotificationChannel";
import { Scan } from "../model/Scan";

export class SlackNotifier extends AbstractWebHookNotifier {
  protected type: NotifierType = NotifierType.SLACK;

  constructor(
    channel: NotificationChannel,
    scan: Scan,
    findings: Finding[],
    args: Object,
  ) {
    super(channel, scan, findings, args);
  }
}

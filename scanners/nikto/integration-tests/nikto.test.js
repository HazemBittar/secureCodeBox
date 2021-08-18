// SPDX-FileCopyrightText: 2021 iteratec GmbH
//
// SPDX-License-Identifier: Apache-2.0

const { scan } = require("../../../tests/integration/helpers");

jest.retryTimes(3);

test(
  "nikto scan against bodgeit demo-target",
  async () => {
    const { categories, severities, count } = await scan(
      "nikto-bodgeit",
      "nikto",
      [
        "-h",
        "bodgeit.demo-targets.svc",
        "-port",
        "8080",
        "-Tuning",
        "1,2,3,5,7,b",
      ], // See nikto bodgeit example
      90
    );

    expect(count).toBe(6);
    expect(categories).toMatchInlineSnapshot(`
      Object {
        "Identified Software": 1,
        "Nikto Finding": 3,
        "X-Content-Type-Options Header": 1,
        "X-Frame-Options Header": 1,
      }
    `);
    expect(severities).toMatchInlineSnapshot(`
      Object {
        "informational": 5,
        "low": 1,
      }
    `);
  },
  3 * 60 * 1000
);

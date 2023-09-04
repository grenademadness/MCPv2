# Master Control Program v2

A rewrite of the old Master Control Program (https://github.com/adh-partnership/bot)
to be more configurable (aka, not hardcoded), modular, and be able to be used in multiple
guilds instead of running multiple bots, one for each guild.

While this may not happen, it may prove useful for some implementations.

## Installation

TODO

## Configuration

### config.yaml

This configuration defines the basic global configuration settings for the bot.

Example config.yaml:

```yaml
database:
  host: 127.0.0.1
  port: "3306"
  user: root
  password: bot
  database: bot
  auto_migrate: true
discord:
  appid: "1234567"
  token: bot.token.discord
```

### Facility Configuration

These are YAML files generally in `facilities/` and are used to define facility guilds.

An example:
```yaml
# Used for internal identification purposes
facility: ZAN
description: Anchorage ARTCC
# The Guild ID for the facility's discord
discord_id: 123456789012345678
# What name format do we set for users in the guild?
# - first_last (John Doe, John Doe | ATM)
# - first_last_initial (John D., John D. | ATM)
name_format: first_last
# Position table, used for the who's online embedded message
# that will be posted to the positions_channel
positions:
- name: Oceanic
  callsigns:
  - prefix: ["ZAN"]
    suffix: ["FSS"]
- name: Enroute
  callsigns:
    prefix: ["ANC"]
    suffix: ["CTR"]
- name: A11
  callsigns:
    prefix: ["ANC"]
    suffix: ["APP"]
- name: FAI TRACON
  callsigns:
    prefix: ["FAI"]
    suffix: ["APP"]
- name: CAB
  callsigns:
    prefix:
    - ADQ
    - AKN
    - ANC
    - BIG
    - BET
    - EIL
    - EDF
    - ENA
    - FAI
    - FBK
    - FRN
    - JNU
    - LHD
    - MRI
    suffix:
    - TWR
    - GND
    - DEL
# The ID number of the channel to post the who's online message to
# This channel should ideally be setup so only the bot posts to it
positions_channel: 1011814580848177195
# Role assignments
# Users who have a configured role but do not meet the conditions will have the role
# removed.
# Conditions are defined in the `if` block and are evaluated as a logical OR.
# Conditions:
# - has_role: Checks if the user has the specified role in the roster
# - controller_type: Check if the user is a `home` controller, `visitor` controller, or `none`
# - unknown: Checks if the account is linked on the roster
roles:
- id: 123456789012345678 # Role ID
  name: ZAN Senior Staff # Role Name (primarily for logging)
  if: # Conditions evaluated as a logical OR
  - condition: has_role
    value: ATM
  - condition: has_role
    value: DATM
  - condition: has_role
    value: TA
- id: 123456789012345678
  name: ZAN Staff
  if:
  - condition: has_role
    value: EC
  - condition: has_role
    value: FE
  - condition: has_role
    value: WM
- id: 123456789012345678
  name: ZAN Assistance Staff
  if:
  - condition: has_role
    value: facilities
  - condition: has_role
    value: events
  - condition: has_role
    value: web
- id: 123456789012345678
  name: Training Team
  if:
  - condition: has_role
    value: INS
  - condition: has_role
    value: MTR
- id: 123456789012345678
  name: FE Team
  if:
  - condition: has_role
    value: facilities
- id: 123456789012345678
  name: Events Team
  if:
  - condition: has_role
    value: events
- id: 123456789012345678
  name: Instructors
  if:
  - condition: has_role
    value: INS
- id: 123456789012345678
  name: Mentors
  if:
  - condition: has_role
    value: MTR
- id: 123456789012345678
  name: ZAN Members
  if:
  - condition: controller_type
    value: home
- id: 123456789012345678
  name: Visitors
  if:
  - condition: controller_type
    value: visit
- id: 123456789012345678
  name: Pilot
  if:
  - condition: controller_type
    value: none
- id: 123456789012345678
  name: Unverified
  if:
  - condition: unknown
    value: true
# Base URL of the API
# Note that the Bot will use this API to fetch event information for the /event
# slash command as well as the roster. The API, if not ADH Partnership's API
# must support the `/v1/events` `/v1/event/:id` and `/v1/user/all` endpoints
# with the same response format as ADH Partnership's API
#
# ADH Partnership's Swagger Docs can be accessed at https://api.zanartcc.org
api: https://api.zanartcc.org
```

## Running

TODO

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](LICENSE) file for details

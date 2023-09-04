/*
 * Copyright Daniel Hawton
 *
 *  Licensed under the Apache License, Version 2.0 (the "License");
 *  you may not use this file except in compliance with the License.
 *  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 *  Unless required by applicable law or agreed to in writing, software
 *  distributed under the License is distributed on an "AS IS" BASIS,
 *  WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  See the License for the specific language governing permissions and
 *  limitations under the License.
 */

package jobs

func UpdateGuilds() {
	/*
		 @TODO
			log.Debugf("Starting UpdateGuilds job")
			defer log.Debugf("Finished UpdateGuilds job")
			for _, f := range facility.FacCfg {
				updateGuild(f)
			}
	*/
}

/* @TODO
func updateGuild(f *facility.Facility) {
	log.Debugf("Requesting Guild Members for %s", f.Facility)
	err := discord.RequestGuildMembers(f.DiscordID, "", 0, "", false)
	if err != nil {
		log.Errorf("Error requesting guild members for %s: %s", f.Facility, err)
		return
	}
}
*/

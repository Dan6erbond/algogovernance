/*
Copyright © 2022 RaviAnand Mohabir moravrav@gmail.com

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

/*
Package main and its subdirectories contain the Algogovernance API client wrapper and the command line interface.

Algogovernance is a full API wrapper for the Algorand Governance Platform enabling developers to get periods, total staked ALGO, voting topics and options, as well as total votes and other data relevant to Algorand governance.

Additionally the repository includes a CLI allowing users to query commonly used data such as information about the current period, and calculate rewards for a given governor.

Using The CLI

The CLI can be installed from Git, and then ran with the algogovernance command:

	$ go install github.com/Dan6erbond/algogovernance
	$ algogovernance currentPeriod

Certain configuration parameters, such as governor can be specified in a .algogovernance.yml file in order to avoid having to pass them as arguments to the CLI.

Save the file to $HOME/.algogovernance.yml and then run the CLI to check that the config is loaded:

	$ algogovernance cfg
	Governor: 3RYOY2LTPC6GLT3ZYE4LUFGGAEMY7GRENZQO7RFNGK2LGCV77QNASK6C6Y

Making Requests

All the client code is made available in the pkg/client folder to to make requests to the Algorand Governance API and supports features such as pagination, sorting and downloads.

Since the Algorand Governance API is unpermissioned, no API keys are required to make requests. Simply call methods from the client package after importing it like so:

	import (
		"github.com/Dan6erbond/algogovernance/pkg/client"
	)

	func main() {
		activePeriod, _ := client.GetActivePeriod()
		governors, _ := client.GetPeriodGovernors(activePeriod.Slug, "", "", "", "cursor", "", "", "")
		for governors.HasNext() {
			governors, _ = governors.GetNext()
		}
	}

Helpers

The package also comes with helpers in the pkg/helpers folder to perform common tasks, such as calculating rewards for a given period:

	import (
		algoRewards "github.com/Dan6erbond/algogovernance/pkg/rewards"
	)

	func main() {
		address := "3RYOY2LTPC6GLT3ZYE4LUFGGAEMY7GRENZQO7RFNGK2LGCV77QNASK6C6Y"
		rewards, _ := algoRewards.GetRewardsForCurrentPeriod(address)
	}
*/
package main

import "github.com/Dan6erbond/algogovernance/cmd"

func main() {
	cmd.Execute()
}

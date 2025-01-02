# Stealth DCA

> **Note:** With the current version, you will need to edit the code slightly to get it running.  
> Don't worry! We're working on making it pretty and user-friendly for our non-tech friends. Stay tuned! 🚀


Stealth DCA is an open-source app for decentralized dollar-cost averaging (DCA) on the Solana blockchain. It uses Jupiter Aggregator for seamless token swaps, ensures MEV protection by routing transactions through your own validator, and maintains complete privacy by running locally on your machine.

Key Features:
- **Decentralized DCA**: Automate your buys with customizable schedules and amounts.
- **Privacy First**: All operations run locally—your keys and data stay safe.
- **ME Protection**: No frontrunning risks
- **Open Source**: Fully transparent and community-driven.
- **Flexible**: No limitations, you can create any size of purchase
- **Fee Free**: Jupiter has 0.1% platform fee 

Start investing smarter and safer with Stealth DCA. 🚀  

## For Developers 🛠️

We welcome contributions from developers! Whether it's fixing a bug, adding a feature, or improving documentation, your work makes a difference.

### How to Contribute
1. Fork the repository and create a new branch for your feature or fix.
2. Make your changes and ensure the code is clean and well-documented.
3. Submit a Pull Request (PR) describing your work and why it's awesome.
4. Our team will review your PR and, if approved, merge it into the main branch.

### Rewards 💰
To encourage contributions, we reward developers with **USD 50 to USD 1000 worth of FAPE** for their merged contributions! Here's how it works:

1. **Create an Issue**: Before starting work on a feature or fix, create an issue in the repository describing the proposed change.
2. **Reward Assignment**: Our team, led by **@gorillafape**, will review the issue and assign a reward amount based on its impact and complexity.
3. **Build the Feature**: Once approved, start building the feature or fix.
4. **Submit a PR**: After completing the work, submit a Pull Request (PR) referencing the approved issue.
5. **Get Rewarded**: Once the PR is reviewed, approved, and merged, the assigned reward will be sent to the developer who implemented the feature.

Contribute, make a difference, and earn rewards while improving the project! 🚀

We are actively looking for contributions in the following areas and will prioritize rewards for these features:

1. **User-Friendly CLI or Web UI**:
    - Build a simple, intuitive CLI or a web interface to easily manage DCA jobs.
    - Make it accessible for non-technical users to create, edit, and monitor DCA jobs.

2. **Custom Destination Wallets**:
    - Add functionality to specify a destination wallet where purchased tokens can be transferred automatically.

3. **State Management**:
    - Implement reading and writing job states using a dedicated state file (e.g., `{jobname}.state`) OR SQLite using https://bun.uptrace.dev/ 
    - This should track progress, last execution time, and other relevant metadata for each DCA job.


Get involved, improve the project, and earn while you code! 🚀  


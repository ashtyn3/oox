By Ashtyn Morel-Blake

*June 16, 2021 -- June 19, 2021*

With the huge release of Ethereum 2.0 and the anticipation of the
changes to come, I wanted to explore Proof of Stake (PoS). My day of
research into PoS led me to the discovery this system can be expanded
upon. This fundamentally expandable idea allows for adoption across so
many blockchains. PoS and its processes for consensus has inspired me to
tweak it a little and try and make my own blockchain --- Oox (pronounced
\"Oo\"-\"x\").

Oox sits on top of the basic foundation of PoS by using people with more
economic trust to decide the order and validity of blocks. At first this
idea sounded terribly unsecure and terribly centralized having the more
fortunate be able to make such drastic decisions. The PoS maintains
trust in the Consensus through locking part of validators (The person
voting on the block) assets into a vault and holds it from them until
the block was successfully validated. If they are successful, they will
get rewarded with one Ether in the case of Ethereum.

Most of my changes to this system our changes in the inherit trust we
give nodes on the network. When making any type of system that is relied
upon by many people I would rather stick to the safer side and use as
many people as possible to confirm something. Using one or two people to
confirm a system seems like too high of trust to give anyone. Here are
my results after producing a new consensus mechanism based on PoS.

**Mechanism.**

**Block Score (*b*)** is a number representation of the risk if the
block is not validated. The higher the block score the more people need
to double check the validity of the block. Block score is calculated
through the finding the standard deviation of number of coins exchanged
plus the highest exchange.

**Block wide gas price (BWGS** or ***G*)** is the bases of my consensus
mechanism. This refers to the amount of money at stake, which can be
used as a reward for validators. Calculating BWGS is block score divided
by ten.

<img src="https://render.githubusercontent.com/render/math?math=b%20=%20\text{stdev}(T)%20+%20\text{max}(T)"/>
<img src="https://render.githubusercontent.com/render/math?math=G%20=%20\frac{b}{10}"/>

These two equations are the foundation of the consensus. The variable
*b* is the block score, and T is an array of all the transaction amounts
in block. ***G*** is the block wide gas score.

**Validator.**

To become a validator node on the network instead of looking at how much
coins you have other nodes will instead look at how many transactions
you minted. Each block successfully validated by the node increases the
nodes **reputation**, which is equal to all the block scores that that
node has confirmed added together. For example, if a validator node has
validated 3 blocks with the following scores --- 10, 9, and 2 (these
block scores are not realistic) there reputation would be 21. Reputation
can also be increased through the number of coins a validator node has
like the original PoS. Validator nodes should not be confused with
normal nodes they are distinctly different. A validator is represented
by an address (like any other user of the blockchain) in the network. A
node is more of a background process and helps in the exchange of
currency and keeping the ledger up to date.

Validators are rewarded after all secondary validators have confirmed
the original validators work. If there is a discrepancy in the block,
which is not caught by the origin validator the BWGS will be deducted
from the faulty validators current balance. This transaction to reward
or deduct the origin validator should be appended to the next block. In
the case that the block is valid the BWGS is spilt between validators.
The origin validator's reputation will be added too by the block score.

Reputation is a tricky number to store. Instead of trying to find a way
to build this into validators addresses I opted for a solution similar
to finding an addresses balance. The solution to this problem is instead
to store the origin validators public key on the block. So when nodes
are deciding on a new validator all they have to do is look through the
chain and find all of the blocks that the validator has validated and
add the block scores together plus the public keys balance. It is
important to make sure that when creating a block that the block score,
validator address is included in the hashed data to prevent changes in
the future.

**Why.**

As for why I decided to make these changes --- I didn't like the
looseness that came with proof of stake. I also didn't like the sense
that people are betting with the validity of blocks. In some ways you
could call this an entirely different consensus mechanism then PoS but
in a way that was my goal is to build off from PoS. The idea of PoS
still stays true with these changes. In case you forgot some of the
benefits of PoS here are some:

1.  Decreased need for high power machines in network.

2.  Drastically lower energy requirements to participate in the network.

3.  Less ability to centralize network consensus.

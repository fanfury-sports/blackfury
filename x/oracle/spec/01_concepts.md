<!--
order: 1
-->

# Concepts

## Voting Procedure

During each `VotePeriod`, the Oracle module obtains consensus on the exchange rates of denominations specified in `VoteTargets` against USD by requiring all members of the validator set to submit a vote for exchange rates before the end of the interval.

Validators must first pre-commit to a exchange rate, then in the subsequent `VotePeriod` submit and reveal their exchange rate alongside a proof that they had pre-commited at that price. This scheme forces the voter to commit to a submission before knowing the votes of others and thereby reduces centralization and free-rider risk in the Oracle.

* Prevote and Vote

    Let `P_t` be the current time interval of duration defined by `VotePeriod` (currently set to 1 minute) during which validators must submit two messages:

    * A `MsgAggregateExchangeRatePrevote`, containing the SHA256 hash of the exchange rates of denominations with respect to USD. A prevote must be submitted for all different denominations on which to report a USD exchange rate.
    * A `MsgAggregateExchangeRateVote`, containing the salt used to create the hash for the aggreagte prevote submitted in the previous interval `P_t-1`.

* Vote Tally

    At the end of `P_t`, the submitted votes are tallied.

    The submitted salt of each vote is used to verify consistency with the prevote submitted by the validator in `P_t-1`. If the validator has not submitted a prevote, or the SHA256 resulting from the salt does not match the hash from the prevote, the vote is dropped.

    For each denomination, if the total voting power of submitted votes exceeds `VoteThreshold` (currently set to 50%), the weighted median of the votes is recorded on-chain as the effective exchange rate for that denomination against USD for the following `VotePeriod` `P_t+1`.

    Denominations receiving fewer than `VoteThreshold` total voting power have their exchange rates deleted from the store, and no swaps can be made with it during the next VotePeriod `P_t+1`.

* Ballot Rewards

    After the votes are tallied, the winners of the ballots are determined with `Tally()`.

    Voters that have managed to vote within a narrow band around the weighted median, are rewarded with a portion of the collected seigniorage. See `k.RewardBallotWinners()` for more details.

## Reward Band

Let `M` be the weighted median, `𝜎` be the standard deviation of the votes in the ballot, and `R` be the `RewardBand` parameter (currently set to 2%), then the band around the median is set to be `𝜀 = max(𝜎, R/2)`. All valid (i.e. bonded and non-jailed) validators that submitted an exchange rate vote in the interval `[M - 𝜀, M + 𝜀]` should be included in the set of winners, weighted by their relative vote power.

## Slashing

> Be sure to read this section carefully as it concerns potential loss of funds.

A `VotePeriod` during which either of the following events occur is considered a "miss":

* The validator fails to submits a vote for exchange rate against USD for **each and every** denomination specified in `VoteTargtes`.

* The validator fails to vote within the `reward band` around the weighted median for one or more denominations.

During every `SlashWindow`(currently set to 1 week), participating validators must maintain a valid vote rate of at least `MinValidPerWindow` (currently set to 5%), lest they get their stake slashed at `SlashFraction`(currently set to 0.01%). The slashed validator is automatically temporarily "jailed" by the protocol (to protect the funds of delegators), and the operator is expected to fix the discrepancy promptly to resume validator participation.

## Abstaining from Voting

A validator may abstain from voting by submitting a non-positive integers for the `ExchangeRates` field in `MsgAggregateExchangeRateVote`. Doing so will absolve them of any penalties for missing `VotePeriod`s, but also disqualify them from receiving Oracle seigniorage rewards for faithful reporting.

## Transitions

The control flow for vote-tallying, exchange rate updates, ballot rewards and slashing happens at the end of every `VotePeriod`, and is found at the [end-block ABCI](./03_end_block.md) function rather than inside message handlers.

# Basic Structure of Blockchain
   In this example we familiarize ourselves with the basic structure of blockchain and terminologies used <p>
   We use nonce to demonstrate how can generation of hash be cpu consuming and demonstrating basic idea of Proof of Work <p>
   Blockchain is stored in a simple array <p>
   We follow TDD in the whole exercise <p>
   If you don't want to look into persistence and is just interested in seeing a simple blockchain, checkout release_v1 and don't bother to read further <p>

# Persistence
   In the second part, we persist our blockchain in BoltDB <p>
   To run this, you need to first get bolt db using "go get github.com/boltdb/bolt" <p>
   Serialization require fields to be exportable, hence all fields are capitalize <p>
   Bolt store key/value pairs in order, since hash wasn't guarantering ordering, used another field to store based on index <p>

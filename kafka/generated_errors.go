package kafka
// Copyright 2016 Confluent Inc.
// AUTOMATICALLY GENERATED BY /home/maglun/gocode/bin/go_rdkafka_generr ON 2018-04-17 10:35:01.163697617 +0200 CEST USING librdkafka 0.11.4-34-gc89764-dirty-devel-O0

/*
#include <librdkafka/rdkafka.h>
*/
import "C"

// ErrorCode is the integer representation of local and broker error codes
type ErrorCode int

// String returns a human readable representation of an error code
func (c ErrorCode) String() string {
      return C.GoString(C.rd_kafka_err2str(C.rd_kafka_resp_err_t(c)))
}

const (
    // ErrBadMsg Local: Bad message format
    ErrBadMsg ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__BAD_MSG)
    // ErrBadCompression Local: Invalid compressed data
    ErrBadCompression ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__BAD_COMPRESSION)
    // ErrDestroy Local: Broker handle destroyed
    ErrDestroy ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__DESTROY)
    // ErrFail Local: Communication failure with broker
    ErrFail ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__FAIL)
    // ErrTransport Local: Broker transport failure
    ErrTransport ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__TRANSPORT)
    // ErrCritSysResource Local: Critical system resource failure
    ErrCritSysResource ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__CRIT_SYS_RESOURCE)
    // ErrResolve Local: Host resolution failure
    ErrResolve ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__RESOLVE)
    // ErrMsgTimedOut Local: Message timed out
    ErrMsgTimedOut ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__MSG_TIMED_OUT)
    // ErrPartitionEOF Broker: No more messages
    ErrPartitionEOF ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__PARTITION_EOF)
    // ErrUnknownPartition Local: Unknown partition
    ErrUnknownPartition ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__UNKNOWN_PARTITION)
    // ErrFs Local: File or filesystem error
    ErrFs ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__FS)
    // ErrUnknownTopic Local: Unknown topic
    ErrUnknownTopic ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__UNKNOWN_TOPIC)
    // ErrAllBrokersDown Local: All broker connections are down
    ErrAllBrokersDown ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__ALL_BROKERS_DOWN)
    // ErrInvalidArg Local: Invalid argument or configuration
    ErrInvalidArg ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__INVALID_ARG)
    // ErrTimedOut Local: Timed out
    ErrTimedOut ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__TIMED_OUT)
    // ErrQueueFull Local: Queue full
    ErrQueueFull ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__QUEUE_FULL)
    // ErrIsrInsuff Local: ISR count insufficient
    ErrIsrInsuff ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__ISR_INSUFF)
    // ErrNodeUpdate Local: Broker node update
    ErrNodeUpdate ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__NODE_UPDATE)
    // ErrSsl Local: SSL error
    ErrSsl ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__SSL)
    // ErrWaitCoord Local: Waiting for coordinator
    ErrWaitCoord ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__WAIT_COORD)
    // ErrUnknownGroup Local: Unknown group
    ErrUnknownGroup ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__UNKNOWN_GROUP)
    // ErrInProgress Local: Operation in progress
    ErrInProgress ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__IN_PROGRESS)
    // ErrPrevInProgress Local: Previous operation in progress
    ErrPrevInProgress ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__PREV_IN_PROGRESS)
    // ErrExistingSubscription Local: Existing subscription
    ErrExistingSubscription ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__EXISTING_SUBSCRIPTION)
    // ErrAssignPartitions Local: Assign partitions
    ErrAssignPartitions ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__ASSIGN_PARTITIONS)
    // ErrRevokePartitions Local: Revoke partitions
    ErrRevokePartitions ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__REVOKE_PARTITIONS)
    // ErrConflict Local: Conflicting use
    ErrConflict ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__CONFLICT)
    // ErrState Local: Erroneous state
    ErrState ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__STATE)
    // ErrUnknownProtocol Local: Unknown protocol
    ErrUnknownProtocol ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__UNKNOWN_PROTOCOL)
    // ErrNotImplemented Local: Not implemented
    ErrNotImplemented ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__NOT_IMPLEMENTED)
    // ErrAuthentication Local: Authentication failure
    ErrAuthentication ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__AUTHENTICATION)
    // ErrNoOffset Local: No offset stored
    ErrNoOffset ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__NO_OFFSET)
    // ErrOutdated Local: Outdated
    ErrOutdated ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__OUTDATED)
    // ErrTimedOutQueue Local: Timed out in queue
    ErrTimedOutQueue ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__TIMED_OUT_QUEUE)
    // ErrUnsupportedFeature Local: Required feature not supported by broker
    ErrUnsupportedFeature ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__UNSUPPORTED_FEATURE)
    // ErrWaitCache Local: Awaiting cache update
    ErrWaitCache ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__WAIT_CACHE)
    // ErrIntr Local: Operation interrupted
    ErrIntr ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__INTR)
    // ErrKeySerialization Local: Key serialization error
    ErrKeySerialization ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__KEY_SERIALIZATION)
    // ErrValueSerialization Local: Value serialization error
    ErrValueSerialization ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__VALUE_SERIALIZATION)
    // ErrKeyDeserialization Local: Key deserialization error
    ErrKeyDeserialization ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__KEY_DESERIALIZATION)
    // ErrValueDeserialization Local: Value deserialization error
    ErrValueDeserialization ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__VALUE_DESERIALIZATION)
    // ErrPartial Local: Partial response
    ErrPartial ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__PARTIAL)
    // ErrReadOnly Local: Read-only object
    ErrReadOnly ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__READ_ONLY)
    // ErrNoent Local: No such entry
    ErrNoent ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__NOENT)
    // ErrUnderflow Local: Read underflow
    ErrUnderflow ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__UNDERFLOW)
    // ErrInvalidType Local: Invalid type
    ErrInvalidType ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR__INVALID_TYPE)
    // ErrUnknown Unknown broker error
    ErrUnknown ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_UNKNOWN)
    // ErrNoError Success
    ErrNoError ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_NO_ERROR)
    // ErrOffsetOutOfRange Broker: Offset out of range
    ErrOffsetOutOfRange ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_OFFSET_OUT_OF_RANGE)
    // ErrInvalidMsg Broker: Invalid message
    ErrInvalidMsg ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_MSG)
    // ErrUnknownTopicOrPart Broker: Unknown topic or partition
    ErrUnknownTopicOrPart ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_UNKNOWN_TOPIC_OR_PART)
    // ErrInvalidMsgSize Broker: Invalid message size
    ErrInvalidMsgSize ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_MSG_SIZE)
    // ErrLeaderNotAvailable Broker: Leader not available
    ErrLeaderNotAvailable ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_LEADER_NOT_AVAILABLE)
    // ErrNotLeaderForPartition Broker: Not leader for partition
    ErrNotLeaderForPartition ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_NOT_LEADER_FOR_PARTITION)
    // ErrRequestTimedOut Broker: Request timed out
    ErrRequestTimedOut ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_REQUEST_TIMED_OUT)
    // ErrBrokerNotAvailable Broker: Broker not available
    ErrBrokerNotAvailable ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_BROKER_NOT_AVAILABLE)
    // ErrReplicaNotAvailable Broker: Replica not available
    ErrReplicaNotAvailable ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_REPLICA_NOT_AVAILABLE)
    // ErrMsgSizeTooLarge Broker: Message size too large
    ErrMsgSizeTooLarge ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_MSG_SIZE_TOO_LARGE)
    // ErrStaleCtrlEpoch Broker: StaleControllerEpochCode
    ErrStaleCtrlEpoch ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_STALE_CTRL_EPOCH)
    // ErrOffsetMetadataTooLarge Broker: Offset metadata string too large
    ErrOffsetMetadataTooLarge ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_OFFSET_METADATA_TOO_LARGE)
    // ErrNetworkException Broker: Broker disconnected before response received
    ErrNetworkException ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_NETWORK_EXCEPTION)
    // ErrGroupLoadInProgress Broker: Group coordinator load in progress
    ErrGroupLoadInProgress ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_GROUP_LOAD_IN_PROGRESS)
    // ErrGroupCoordinatorNotAvailable Broker: Group coordinator not available
    ErrGroupCoordinatorNotAvailable ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_GROUP_COORDINATOR_NOT_AVAILABLE)
    // ErrNotCoordinatorForGroup Broker: Not coordinator for group
    ErrNotCoordinatorForGroup ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_NOT_COORDINATOR_FOR_GROUP)
    // ErrTopicException Broker: Invalid topic
    ErrTopicException ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_TOPIC_EXCEPTION)
    // ErrRecordListTooLarge Broker: Message batch larger than configured server segment size
    ErrRecordListTooLarge ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_RECORD_LIST_TOO_LARGE)
    // ErrNotEnoughReplicas Broker: Not enough in-sync replicas
    ErrNotEnoughReplicas ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_NOT_ENOUGH_REPLICAS)
    // ErrNotEnoughReplicasAfterAppend Broker: Message(s) written to insufficient number of in-sync replicas
    ErrNotEnoughReplicasAfterAppend ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_NOT_ENOUGH_REPLICAS_AFTER_APPEND)
    // ErrInvalidRequiredAcks Broker: Invalid required acks value
    ErrInvalidRequiredAcks ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_REQUIRED_ACKS)
    // ErrIllegalGeneration Broker: Specified group generation id is not valid
    ErrIllegalGeneration ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_ILLEGAL_GENERATION)
    // ErrInconsistentGroupProtocol Broker: Inconsistent group protocol
    ErrInconsistentGroupProtocol ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INCONSISTENT_GROUP_PROTOCOL)
    // ErrInvalidGroupID Broker: Invalid group.id
    ErrInvalidGroupID ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_GROUP_ID)
    // ErrUnknownMemberID Broker: Unknown member
    ErrUnknownMemberID ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_UNKNOWN_MEMBER_ID)
    // ErrInvalidSessionTimeout Broker: Invalid session timeout
    ErrInvalidSessionTimeout ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_SESSION_TIMEOUT)
    // ErrRebalanceInProgress Broker: Group rebalance in progress
    ErrRebalanceInProgress ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_REBALANCE_IN_PROGRESS)
    // ErrInvalidCommitOffsetSize Broker: Commit offset data size is not valid
    ErrInvalidCommitOffsetSize ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_COMMIT_OFFSET_SIZE)
    // ErrTopicAuthorizationFailed Broker: Topic authorization failed
    ErrTopicAuthorizationFailed ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_TOPIC_AUTHORIZATION_FAILED)
    // ErrGroupAuthorizationFailed Broker: Group authorization failed
    ErrGroupAuthorizationFailed ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_GROUP_AUTHORIZATION_FAILED)
    // ErrClusterAuthorizationFailed Broker: Cluster authorization failed
    ErrClusterAuthorizationFailed ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_CLUSTER_AUTHORIZATION_FAILED)
    // ErrInvalidTimestamp Broker: Invalid timestamp
    ErrInvalidTimestamp ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_TIMESTAMP)
    // ErrUnsupportedSaslMechanism Broker: Unsupported SASL mechanism
    ErrUnsupportedSaslMechanism ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_UNSUPPORTED_SASL_MECHANISM)
    // ErrIllegalSaslState Broker: Request not valid in current SASL state
    ErrIllegalSaslState ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_ILLEGAL_SASL_STATE)
    // ErrUnsupportedVersion Broker: API version not supported
    ErrUnsupportedVersion ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_UNSUPPORTED_VERSION)
    // ErrTopicAlreadyExists Broker: Topic already exists
    ErrTopicAlreadyExists ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_TOPIC_ALREADY_EXISTS)
    // ErrInvalidPartitions Broker: Invalid number of partitions
    ErrInvalidPartitions ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_PARTITIONS)
    // ErrInvalidReplicationFactor Broker: Invalid replication factor
    ErrInvalidReplicationFactor ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_REPLICATION_FACTOR)
    // ErrInvalidReplicaAssignment Broker: Invalid replica assignment
    ErrInvalidReplicaAssignment ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_REPLICA_ASSIGNMENT)
    // ErrInvalidConfig Broker: Configuration is invalid
    ErrInvalidConfig ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_CONFIG)
    // ErrNotController Broker: Not controller for cluster
    ErrNotController ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_NOT_CONTROLLER)
    // ErrInvalidRequest Broker: Invalid request
    ErrInvalidRequest ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_REQUEST)
    // ErrUnsupportedForMessageFormat Broker: Message format on broker does not support request
    ErrUnsupportedForMessageFormat ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_UNSUPPORTED_FOR_MESSAGE_FORMAT)
    // ErrPolicyViolation Broker: Isolation policy volation
    ErrPolicyViolation ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_POLICY_VIOLATION)
    // ErrOutOfOrderSequenceNumber Broker: Broker received an out of order sequence number
    ErrOutOfOrderSequenceNumber ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_OUT_OF_ORDER_SEQUENCE_NUMBER)
    // ErrDuplicateSequenceNumber Broker: Broker received a duplicate sequence number
    ErrDuplicateSequenceNumber ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_DUPLICATE_SEQUENCE_NUMBER)
    // ErrInvalidProducerEpoch Broker: Producer attempted an operation with an old epoch
    ErrInvalidProducerEpoch ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_PRODUCER_EPOCH)
    // ErrInvalidTxnState Broker: Producer attempted a transactional operation in an invalid state
    ErrInvalidTxnState ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_TXN_STATE)
    // ErrInvalidProducerIDMapping Broker: Producer attempted to use a producer id which is not currently assigned to its transactional id
    ErrInvalidProducerIDMapping ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_PRODUCER_ID_MAPPING)
    // ErrInvalidTransactionTimeout Broker: Transaction timeout is larger than the maximum value allowed by the broker's max.transaction.timeout.ms
    ErrInvalidTransactionTimeout ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_INVALID_TRANSACTION_TIMEOUT)
    // ErrConcurrentTransactions Broker: Producer attempted to update a transaction while another concurrent operation on the same transaction was ongoing
    ErrConcurrentTransactions ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_CONCURRENT_TRANSACTIONS)
    // ErrTransactionCoordinatorFenced Broker: Indicates that the transaction coordinator sending a WriteTxnMarker is no longer the current coordinator for a given producer
    ErrTransactionCoordinatorFenced ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_TRANSACTION_COORDINATOR_FENCED)
    // ErrTransactionalIDAuthorizationFailed Broker: Transactional Id authorization failed
    ErrTransactionalIDAuthorizationFailed ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_TRANSACTIONAL_ID_AUTHORIZATION_FAILED)
    // ErrSecurityDisabled Broker: Security features are disabled
    ErrSecurityDisabled ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_SECURITY_DISABLED)
    // ErrOperationNotAttempted Broker: Operation not attempted
    ErrOperationNotAttempted ErrorCode = ErrorCode(C.RD_KAFKA_RESP_ERR_OPERATION_NOT_ATTEMPTED)
)

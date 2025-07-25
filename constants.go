/*
 * MinIO Go Library for Amazon S3 Compatible Cloud Storage
 * Copyright 2015-2017 MinIO, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package ossClient

// Multipart upload defaults.

// absMinPartSize - absolute minimum part size (5 MiB) below which
// a part in a multipart upload may not be uploaded.
const absMinPartSize = 1024 * 1024 * 5

// minPartSize - minimum part size 16MiB per object after which
// putObject behaves internally as multipart.
const minPartSize = 1024 * 1024 * 16

// maxPartsCount - maximum number of parts for a single multipart session.
const maxPartsCount = 10000

// maxPartSize - maximum part size 5GiB for a single multipart upload
// operation.
const maxPartSize = 1024 * 1024 * 1024 * 5

// maxSinglePutObjectSize - maximum size 5GiB of object per PUT
// operation.
const maxSinglePutObjectSize = 1024 * 1024 * 1024 * 5

// maxMultipartPutObjectSize - maximum size 5TiB of object for
// Multipart operation.
const maxMultipartPutObjectSize = 1024 * 1024 * 1024 * 1024 * 5

// unsignedPayload - value to be set to X-Amz-Content-Sha256 header when
// we don't want to sign the request payload
const unsignedPayload = "UNSIGNED-PAYLOAD"

// unsignedPayloadTrailer value to be set to X-Amz-Content-Sha256 header when
// we don't want to sign the request payload, but have a trailer.
const unsignedPayloadTrailer = "STREAMING-UNSIGNED-PAYLOAD-TRAILER"

// Total number of parallel workers used for multipart operation.
const totalWorkers = 4

// Signature related constants.
const (
	signV4Algorithm   = "AWS4-HMAC-SHA256"
	iso8601DateFormat = "20060102T150405Z"
)

const (
	// Storage class header.
	amzStorageClass = "X-Amz-Storage-Class"

	// Website redirect location header
	amzWebsiteRedirectLocation = "X-Amz-Website-Redirect-Location"

	// Object Tagging headers
	amzTaggingHeader          = "X-Amz-Tagging"
	amzTaggingHeaderDirective = "X-Amz-Tagging-Directive"

	amzVersionID         = "X-Amz-Version-Id"
	amzTaggingCount      = "X-Amz-Tagging-Count"
	amzExpiration        = "X-Amz-Expiration"
	amzRestore           = "X-Amz-Restore"
	amzReplicationStatus = "X-Amz-Replication-Status"
	amzDeleteMarker      = "X-Amz-Delete-Marker"

	// Object legal hold header
	amzLegalHoldHeader = "X-Amz-Object-Lock-Legal-Hold"

	// Object retention header
	amzLockMode         = "X-Amz-Object-Lock-Mode"
	amzLockRetainUntil  = "X-Amz-Object-Lock-Retain-Until-Date"
	amzBypassGovernance = "X-Amz-Bypass-Governance-Retention"

	// Replication status
	amzBucketReplicationStatus = "X-Amz-Replication-Status"
	// Minio specific Replication/lifecycle transition extension
	minIOBucketSourceMTime = "X-Minio-Source-Mtime"

	minIOBucketSourceETag              = "X-Minio-Source-Etag"
	minIOBucketReplicationDeleteMarker = "X-Minio-Source-DeleteMarker"
	minIOBucketReplicationProxyRequest = "X-Minio-Source-Proxy-Request"
	minIOBucketReplicationRequest      = "X-Minio-Source-Replication-Request"
	minIOBucketReplicationCheck        = "X-Minio-Source-Replication-Check"

	// Header indicates last tag update time on source
	minIOBucketReplicationTaggingTimestamp = "X-Minio-Source-Replication-Tagging-Timestamp"
	// Header indicates last retention update time on source
	minIOBucketReplicationObjectRetentionTimestamp = "X-Minio-Source-Replication-Retention-Timestamp"
	// Header indicates last legalhold update time on source
	minIOBucketReplicationObjectLegalHoldTimestamp = "X-Minio-Source-Replication-LegalHold-Timestamp"
	minIOForceDelete                               = "x-minio-force-delete"
	// Header indicates delete marker replication request can be sent by source now.
	minioTgtReplicationReady = "X-Minio-Replication-Ready"
	// Header asks if delete marker replication request can be sent by source now.
	isMinioTgtReplicationReady = "X-Minio-Check-Replication-Ready"
	/* trinet */
	MinIOPartialUpdateMode       = "X-Minio-Partial-Update-Mode"
	MinIOPartialUpdateOffset     = "X-Minio-Partial-Update-Offset"
	AmzSnowballExtract           = "X-Amz-Meta-Snowball-Auto-Extract"
	MinIOSnowballIgnoreDirs      = "X-Amz-Meta-Minio-Snowball-Ignore-Dirs"
	MinIOMergeMultipart          = "X-Minio-Merge-Multipart"
	MinIOPoolEngine              = "X-Minio-Pool-Engine"
	MinIODelBucketParallelDrives = "X-Minio-Del-Bucket-Parallel-Drives"
	MinIODelPrefixParallelDrives = "X-Minio-Del-Prefix-Parallel-Drives"
	MinIOSnowballUpdateMTime     = "X-Amz-Meta-Minio-Snowball-Update-MTime"
	/* trinet */
)

package ethdb

import (
	"context"
	"errors"
	"github.com/ledgerwatch/lmdb-go/lmdb"
	"github.com/ledgerwatch/turbo-geth/common/dbutils"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestBucketCRUD(t *testing.T) {
	require := require.New(t)
	kv := NewLMDB().InMem().MustOpen()
	defer kv.Close()

	ctx := context.Background()
	if err := kv.Update(ctx, func(tx Tx) error {
		normalBucket := dbutils.Buckets[15]
		deprecatedBucket := dbutils.DeprecatedBuckets[0]
		migrator, ok := tx.(BucketMigrator)
		if !ok {
			return nil
		}

		// check thad buckets have unique DBI's
		uniquness := map[lmdb.DBI]bool{}
		castedKv := kv.(*LmdbKV)
		for _, dbi := range castedKv.buckets {
			if dbi == NonExistingDBI {
				continue
			}
			_, ok := uniquness[dbi]
			require.False(ok)
			uniquness[dbi] = true
		}

		require.True(migrator.ExistsBucket(normalBucket))
		require.True(errors.Is(migrator.DropBucket(normalBucket), ErrAttemptToDeleteNonDeprecatedBucket))

		require.False(migrator.ExistsBucket(deprecatedBucket))
		require.NoError(migrator.CreateBucket(deprecatedBucket))
		require.True(migrator.ExistsBucket(deprecatedBucket))

		require.NoError(migrator.DropBucket(deprecatedBucket))
		require.False(migrator.ExistsBucket(deprecatedBucket))

		require.NoError(migrator.CreateBucket(deprecatedBucket))
		require.True(migrator.ExistsBucket(deprecatedBucket))

		err := tx.Cursor(deprecatedBucket).Put([]byte{1}, []byte{1})
		require.NoError(err)
		v, err := tx.Get(deprecatedBucket, []byte{1})
		require.NoError(err)
		require.Equal([]byte{1}, v)

		buckets, err := migrator.ExistingBuckets()
		require.NoError(err)
		require.True(len(buckets) > 10)

		// check thad buckets have unique DBI's
		uniquness = map[lmdb.DBI]bool{}
		for _, dbi := range castedKv.buckets {
			if dbi == NonExistingDBI {
				continue
			}
			_, ok := uniquness[dbi]
			require.False(ok)
			uniquness[dbi] = true
		}
		return nil
	}); err != nil {
		require.NoError(err)
	}
}

package goplayground

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

var m map[string]int

func BenchmarkMap(b *testing.B) {
	length := 65
	keys := make([]string, length)
	for k := 0; k < length; k++ {
		keys[k] = fmt.Sprintf("key%d", k+1)
	}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m = make(map[string]int, length)
		for k := 0; k < length; k++ {
			m[keys[k]] = rand.Int()
		}
	}
	b.StopTimer()
	assert.Equal(b, length, len(m), "Ensure there are no duplicate keys")
	mapLiteral := buildMapUsingLiteral65()
	for k := range mapLiteral {
		_, ok := m[k]
		assert.Equal(b, true, ok, "Ensure %v is present", k)
	}
}

func BenchmarkMapStringLiteral(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		m = buildMapUsingLiteral65()
	}
}

func buildMapUsingLiteral65() map[string]int {
	return map[string]int{
		"key1":  rand.Int(),
		"key2":  rand.Int(),
		"key3":  rand.Int(),
		"key4":  rand.Int(),
		"key5":  rand.Int(),
		"key6":  rand.Int(),
		"key7":  rand.Int(),
		"key8":  rand.Int(),
		"key9":  rand.Int(),
		"key10": rand.Int(),
		"key11": rand.Int(),
		"key12": rand.Int(),
		"key13": rand.Int(),
		"key14": rand.Int(),
		"key15": rand.Int(),
		"key16": rand.Int(),
		"key17": rand.Int(),
		"key18": rand.Int(),
		"key19": rand.Int(),
		"key20": rand.Int(),
		"key21": rand.Int(),
		"key22": rand.Int(),
		"key23": rand.Int(),
		"key24": rand.Int(),
		"key25": rand.Int(),
		"key26": rand.Int(),
		"key27": rand.Int(),
		"key28": rand.Int(),
		"key29": rand.Int(),
		"key30": rand.Int(),
		"key31": rand.Int(),
		"key32": rand.Int(),
		"key33": rand.Int(),
		"key34": rand.Int(),
		"key35": rand.Int(),
		"key36": rand.Int(),
		"key37": rand.Int(),
		"key38": rand.Int(),
		"key39": rand.Int(),
		"key40": rand.Int(),
		"key41": rand.Int(),
		"key42": rand.Int(),
		"key43": rand.Int(),
		"key44": rand.Int(),
		"key45": rand.Int(),
		"key46": rand.Int(),
		"key47": rand.Int(),
		"key48": rand.Int(),
		"key49": rand.Int(),
		"key50": rand.Int(),
		"key51": rand.Int(),
		"key52": rand.Int(),
		"key53": rand.Int(),
		"key54": rand.Int(),
		"key55": rand.Int(),
		"key56": rand.Int(),
		"key57": rand.Int(),
		"key58": rand.Int(),
		"key59": rand.Int(),
		"key60": rand.Int(),
		"key61": rand.Int(),
		"key62": rand.Int(),
		"key63": rand.Int(),
		"key64": rand.Int(),
		"key65": rand.Int(),
	}
}

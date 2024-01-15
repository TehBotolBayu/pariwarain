package com.example.pariwarain.ui.properties.posting

import androidx.appcompat.app.AppCompatActivity
import android.os.Bundle
import androidx.recyclerview.widget.LinearLayoutManager
import com.example.pariwarain.R
import com.example.pariwarain.adapter.CategoryAdapter
import com.example.pariwarain.databinding.ActivityPostingBinding

class PostingActivity : AppCompatActivity() {

    private lateinit var binding: ActivityPostingBinding

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        binding = ActivityPostingBinding.inflate(layoutInflater)
        setContentView(binding.root)
        setupToolbar()

        val listTag = listOf("Rumah", "Gedung", "Toko", "Billboard", "Kantor", "Cafe")
        binding.rvCategory.setHasFixedSize(true)
        binding.rvCategory.layoutManager = LinearLayoutManager(this, LinearLayoutManager.HORIZONTAL, false)
        val adapter = CategoryAdapter(listTag)
//        val activeCategories = adapter.getListCategoryActive()
        binding.rvCategory.adapter = adapter
//        Log.d("listCategoryActiveMainActivity", "$activeCategories | Size => ${activeCategories.size}")
    }

    private fun setupToolbar() {
        val toolbar = binding.toolbar

        toolbar.setNavigationOnClickListener {
            super.onBackPressed()
        }
    }
}
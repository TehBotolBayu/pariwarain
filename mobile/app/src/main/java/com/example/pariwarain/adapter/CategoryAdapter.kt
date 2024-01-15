package com.example.pariwarain.adapter

import android.util.Log
import android.view.LayoutInflater
import android.view.ViewGroup
import android.widget.Toast
import androidx.core.content.ContextCompat
import androidx.recyclerview.widget.RecyclerView
import com.example.pariwarain.R
import com.example.pariwarain.databinding.CategoryItemBinding

class CategoryAdapter(private val listCategories: List<String>) : RecyclerView.Adapter<CategoryAdapter.CategoryViewHolder>() {

    private var listCategoryActive = mutableListOf<String>()

    override fun onCreateViewHolder(parent: ViewGroup, viewType: Int): CategoryViewHolder {
        val binding = CategoryItemBinding.inflate(LayoutInflater.from(parent.context), parent, false)
        return CategoryViewHolder(binding)
    }

    override fun onBindViewHolder(holder: CategoryViewHolder, position: Int) {
        val category = listCategories[position]
        holder.bind(category)
    }

    override fun getItemCount(): Int {
        return listCategories.size
    }

    override fun getItemViewType(position: Int): Int {
        return when (position) {
            0 -> ViewType.FIRST_ITEM.ordinal
            listCategories.size - 1 -> ViewType.LAST_ITEM.ordinal
            else -> ViewType.OTHER_ITEMS.ordinal
        }
    }

    fun getListCategoryActive(): List<String> {
        return listCategoryActive
    }

    inner class CategoryViewHolder(private val binding: CategoryItemBinding) : RecyclerView.ViewHolder(binding.root) {
        var isSelected = false
        fun bind(category: String) {
            binding.tag1.text = category
            binding.containerTag1.setOnClickListener { item ->
                Toast.makeText(item.context, category, Toast.LENGTH_SHORT).show()
                isSelected = !isSelected
                if (isSelected) {
                    if (category !in listCategoryActive) {
                        listCategoryActive.add(category)
                    }
                    binding.tag1.setTextColor(ContextCompat.getColor(item.context, R.color.tag_on))
                    binding.containerTag1.setBackgroundResource(R.drawable.tag_on)
                } else {
                    if (category in listCategoryActive) {
                        listCategoryActive.remove(category)
                    }
                    binding.tag1.setTextColor(ContextCompat.getColor(item.context, R.color.tag_off))
                    binding.containerTag1.setBackgroundResource(R.drawable.tag_off)
                }
                Log.d("listCategoryActive", "$listCategoryActive | Size => ${listCategoryActive.size}")
            }

            // set margin for first and last item, and then margin for separator each item
            val marginStartEnd = 0
            val marginSeparator = 4
            val density = binding.root.context.resources.displayMetrics.density
            val startEnd = (marginStartEnd * density).toInt()
            val separator = (marginSeparator * density).toInt()
            val layoutParams = binding.root.layoutParams as RecyclerView.LayoutParams

            when (getItemViewType(adapterPosition)) {
                ViewType.FIRST_ITEM.ordinal -> {
                    layoutParams.marginStart = startEnd
                    layoutParams.marginEnd = separator
                }
                ViewType.LAST_ITEM.ordinal -> {
                    layoutParams.marginStart = separator
                    layoutParams.marginEnd = startEnd
                }
                else -> {
                    layoutParams.marginStart = separator
                    layoutParams.marginEnd = separator
                }
            }

            binding.root.layoutParams = layoutParams
        }
    }

    enum class ViewType {
        FIRST_ITEM,
        LAST_ITEM,
        OTHER_ITEMS
    }
}
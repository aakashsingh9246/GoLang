function mergeInversion(arr){
	let sortedArray = mergeSort(arr);
	let indexArray = []
	sortedArray.forEach((val) => {
		
		indexArray.push(arr.indexOf(val));
	});
	console.log(indexArray);
	let count = minimumSwaps(indexArray);
	return count;



	function mergeSort(arr){
		if (arr.length<=1){
			return arr;
		}
		const middle = Math.floor(arr.length/2);
		const left = arr.slice(0,middle);
		const right = arr.slice(middle);
		return merge(mergeSort(left), mergeSort(right));
	}

function merge (left, right){
		let merged = [], leftIndex = 0, rightIndex=0;
		while (leftIndex<left.length && rightIndex<right.length){
			if (left[leftIndex]<right[rightIndex]){
				merged.push(left[leftIndex]);
				leftIndex++
			}else{
				merged.push(right[rightIndex]);
				rightIndex++
			}
		};
		return merged.concat(left.slice(leftIndex)).concat(right.slice(rightIndex));
	}

function minimumSwaps(arr){
		let swap = 0;
		let boolArray = []
		arr.forEach((val,i)=>{
			let j = i;
			let cycle = 0;
			while (!boolArray[j]){
				boolArray[j] = true;
				j = (arr[j]-1);
				cycle++;
			}
			if (cycle!=0){
				swap+=(cycle-1);
			}
		});
		return swap
	}
}





console.log(mergeInversion([1,1,1,2,2]));


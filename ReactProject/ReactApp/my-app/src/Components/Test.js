import React, { Component } from 'react';
import { Grommet, Heading, Button, Box } from 'grommet'

export default class Test extends Component {
	constructor(props){
		super(props)
		
		this.createFile = this.createFile.bind(this)
		this.updateFile = this.updateFile.bind(this)
		this.deleteFile = this.deleteFile.bind(this)
		this.state ={
			val: undefined
		}
	}

	createFile(){
		var xhr = new XMLHttpRequest();
		var temp = 0	
        xhr.open('POST', 'http://localhost:8080/start', false);
        xhr.addEventListener('readystatechange', function(){
            if(xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200){
                var item = xhr.responseText;
                console.log(item);
                temp = item
            }
        });
        xhr.send();
        this.setState({
        	val : temp
        })
	}

	updateFile(){
		var xhr = new XMLHttpRequest();
		var temp = 0	
        xhr.open('POST', 'http://localhost:8080/update', false);
        xhr.addEventListener('readystatechange', function(){
            if(xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200){
                var item = xhr.responseText;
                console.log(item);
                temp = item
            }
        });
        xhr.send();
        this.setState({
        	val : temp
        })
	}

	deleteFile(){
		var xhr = new XMLHttpRequest();
		var temp = 0	
        xhr.open('POST', 'http://localhost:8080/delete', false);
        xhr.addEventListener('readystatechange', function(){
            if(xhr.readyState === XMLHttpRequest.DONE && xhr.status === 200){
                var item = xhr.responseText;
                console.log(item);
                temp = item
            }
        });
        xhr.send();
        this.setState({
        	val : temp
        })
	}

	componentDidMount(){
	
	}

	componentWillUnmount(){
		
	}

	render() {
		let timeZone
		if(this.state.val){
			timeZone = (<Box border={{ color: 'brand', size: 'small' }}
				pad="medium"
				margin={"small"}>
				{this.state.val}

			</Box>)
		}

		return (
			<Grommet className="App">
			<Box direction={'row'}>
				<Box>
      				<Button
  						label="Start"
  						onClick={() => {this.createFile()}}/>
  				</Box>
  				<Box>
      				<Button
  						label="Update"
  						onClick={() => {this.updateFile()}}/>
  				</Box>
  				<Box>
      				<Button
  						label="Delete"
  						onClick={() => {this.deleteFile()}}/>
  				</Box>
  				</Box>
  					{timeZone}
			</Grommet>
		)
	}
}
import React, { Component } from 'react';
import { Grommet, Heading, Button, Box } from 'grommet'

export default class Test extends Component {
	constructor(props){
		super(props)
		this.onclick = this.onclick.bind(this)
		this.updateTime = this.updateTime.bind(this)
		this.state ={
			showTime: false,
			time: undefined
		}
	}

	onclick(){
		this.setState({
			showTime: true
		})

	}

	updateTime(){
			this.setState({
				time: new Date().toLocaleString()
			})
	}

	componentDidMount(){
		this.setIntervalID = setInterval(() => this.updateTime(),1000)
	}

	componentWillUnmount(){
		clearInterval(this.setIntervalID)
	}

	render() {
		let timeZone
		if(this.state.showTime){
			timeZone = (<Box border={{ color: 'brand', size: 'small' }}
				pad="medium"
				margin={"small"}>
				{this.state.time}

			</Box>)
		}

		return (
			<Grommet className="App">
      			<Button
  					label="Start"
  					onClick={() => {this.onclick()}}/>
  					{timeZone}
			</Grommet>
		)
	}
}
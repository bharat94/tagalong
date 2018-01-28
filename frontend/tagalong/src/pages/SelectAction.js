import React from 'react';
import { StyleSheet, Text, View } from 'react-native';
import { DrawerNavigator } from 'react-navigation'

import Profile from './src/pages/Profile.js'
import MyEvents from './src/pages/MyEvents.js'
import SelectEvent from './src/pages/SelectEvent.js'

const Navigation = DrawerNavigator(
	{
		"Select Event": {
			screen: SelectEvent
		},
		Profile: {
			screen: Profile
		},
		"My Events": {
			screen: MyEvents
		}
	}
)

export default Navigation;


import React from 'react';

import {id as pluginId} from '../../manifest';

const {formatText, messageHtmlToComponent} = window.PostUtils;

const FAKE_DATA = {
    'smiley': 24
};

export default class UserReputation extends React.PureComponent {
    constructor(props) {
        super(props);
        props.actions.fetchUserReputations(props.userId);
    }
    render() {
        const {userReputations} = this.props;
        console.log('render', userReputations)
        if(userReputations == null || Object.keys(userReputations).length == 0) {
            return <React.Fragment />
        }
        console.log('keys', Object.keys(userReputations))
        return (
            <React.Fragment>
                <hr className="divider divider--expanded" />
                {Object.keys(userReputations).map(function(userReputationKey){
                    const formattedText = formatText(':'+userReputationKey+': ('+userReputations[userReputationKey]+')');
                    return (
                        <div 
                            key={pluginId + '_' + userReputationKey}
                            style={style.row}
                        >
                            {messageHtmlToComponent(formattedText)}
                        </div>
                    )
                })}
            </React.Fragment>
        );
    }
}

const style = {
    row: {
        minHeight: '23px', // TODO - Adjust style for prettiness, and ensure it works on all themes
    },
};